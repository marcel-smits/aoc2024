Param (
    [int]$Day
)

# Define directories and files
$DayFolder = "src/day$Day"
$InputFolder = "$DayFolder/inputs"

# Create folders
New-Item -ItemType Directory -Force -Path $DayFolder, $InputFolder

# Create sample.txt
Out-File -FilePath "$InputFolder/sample.txt" -Encoding utf8

# Create main.go
@"
package main

import (
    "os"
    "fmt"
)

func main() {
    p1("inputs/sample.txt")
    p2("inputs/sample.txt")
}

func p1(path string) int {
    fmt.Print(path)

    lines, err := os.ReadFile(path)

    if err != nil {
        return 0
    }

    return len(lines)
}

func p2(path string) int {
    fmt.Print(path)

    lines, err := os.ReadFile(path)

    if err != nil {
        return 0
    }


    return len(lines)
}

"@ | Out-File -FilePath "$DayFolder/main.go" -Encoding utf8

# Create main_test.go
@"
package main

import "testing"

func Test_p1(t *testing.T) {
    type args struct {
        path string
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "sample",
            args: args{
                    path: "inputs/sample.txt",
                    },
                    want: 0,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := p1(tt.args.path); got != tt.want {
                t.Errorf("p1() %v, want %v.", got, tt.want)
            }
        })
    }
}

func Test_p2(t *testing.T) {
    type args struct {
        path string
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "sample",
            args: args{
                    path: "inputs/sample.txt",
                    },
                    want: 0,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := p2(tt.args.path); got != tt.want {
                t.Errorf("p2() %v, want %v.", got, tt.want)
            }
        })
    }
}
"@ | Out-File -FilePath "$DayFolder/main_test.go" -Encoding utf8

Write-Output "Setup complete for Day $Day."