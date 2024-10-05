package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Variable struct {
    name  string
    typ   string // can be "int" or "string"
    value interface{}
}

var variables = make(map[string]Variable)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: ./EspressoScript <file>")
        return
    }

    fileName := os.Args[1]
    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var programLines []string
    for scanner.Scan() {
        programLines = append(programLines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    execute(programLines)
}

func execute(lines []string) {
    var inFunction bool
    var functionLines []string

    for _, line := range lines {
        line = strings.TrimSpace(line)

        if strings.HasPrefix(line, "let ") {
            defineVariable(line)
        } else if strings.HasPrefix(line, "fun ") {
            inFunction = true
            functionLines = append(functionLines, line)
        } else if inFunction && line == "end" {
            inFunction = false
            executeFunction(functionLines)
            functionLines = nil
        } else if inFunction {
            functionLines = append(functionLines, line)
        }
    }
}

func executeFunction(lines []string) {
    for _, line := range lines {
        line = strings.TrimSpace(line)

        if strings.HasPrefix(line, "if ") {
            evaluateIf(line, lines)
        } else if strings.HasPrefix(line, "printf(") {
            handlePrintf(line)
        }
    }
}

func defineVariable(line string) {
    parts := strings.Split(line, "=")
    nameTyp := strings.Fields(parts[0])[1] // grab variable type
    name := strings.Fields(parts[0])[1]     // grab variable name
    value := strings.TrimSpace(parts[1])

    if nameTyp == "int" {
        intValue, _ := strconv.Atoi(value)
        variables[name] = Variable{name: name, typ: nameTyp, value: intValue}
    } else if nameTyp == "string" {
        variables[name] = Variable{name: name, typ: nameTyp, value: strings.Trim(value, "\"")}
    }
}

func evaluateIf(line string, lines []string) {
    condition := strings.TrimPrefix(line, "if ")
    condition = strings.TrimSpace(condition[:len(condition)-1]) // remove the last character
    parts := strings.Fields(condition)

    if len(parts) != 3 {
        return
    }

    var var1, operator, var2 string
    var1 = parts[0]
    operator = parts[1]
    var2 = parts[2]

    v1 := variables[var1]
    v2 := variables[var2]

    if v1.typ == "int" && v2.typ == "int" {
        switch operator {
        case ">":
            if v1.value.(int) > v2.value.(int) {
                // If condition is true, execute lines until end or else
                for _, innerLine := range lines {
                    innerLine = strings.TrimSpace(innerLine)
                    if innerLine == "else" {
                        break
                    }
                    if innerLine == "end" {
                        return // Stop execution at end
                    }
                    if innerLine != "" {
                        handlePrintf(innerLine)
                    }
                }
            } else {
                // If condition is false, check for the else statement
                for _, innerLine := range lines {
                    innerLine = strings.TrimSpace(innerLine)
                    if innerLine == "end" {
                        return // Stop execution at end
                    }
                    if innerLine == "else" {
                        continue // Skip to else
                    }
                    if innerLine != "" {
                        handlePrintf(innerLine)
                    }
                }
            }
        }
    }
}

func handlePrintf(line string) {
    line = strings.TrimSpace(line)
    line = strings.TrimPrefix(line, "printf(")
    line = strings.TrimSuffix(line, ")")

    // Split arguments by comma
    parts := strings.Split(line, ",")
    var output strings.Builder

    for _, part := range parts {
        part = strings.TrimSpace(part)

        if strings.HasPrefix(part, "\"") && strings.HasSuffix(part, "\"") {
            // String literals
            output.WriteString(strings.Trim(part, "\"")) 
        } else if variable, exists := variables[part]; exists {
            // Variable substitution
            if variable.typ == "string" {
                output.WriteString(variable.value.(string))
            } else if variable.typ == "int" {
                output.WriteString(fmt.Sprintf("%d", variable.value))
            }
        }
        
        // Add space for separation
        output.WriteString(" ")
    }

    fmt.Println(strings.TrimSpace(output.String()))  // Trim white spaces
}
