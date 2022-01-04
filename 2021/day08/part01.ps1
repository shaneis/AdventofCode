#! /usr/bin/pwsh

[CmdletBinding()]
param (
    $Path
)

enum Segments {
    One = 2    
    Four = 4
    Seven = 3
    Eight = 7
}

$Lines = Get-Content -Path $Path

[int64] $count = 0
foreach ($line in $lines) {
    $ls = -split $line
    [bool] $output = $false
    foreach ($l in $ls) {
        $C = if ($l.Length -in (
                [Segments]::One,
                [Segments]::Four,
                [Segments]::Seven,
                [Segments]::Eight
            ) -and $output) {
            'Green'    
            $count++
        } else {
            'Red'
        }

        if (-not $output) {
            $output = $l -eq '|' ? $true : $false
        }

        Write-Host "$l " -foreground $C -NoNewLine
    }
    Write-Host ([Environment]::NewLine)
}
"Part 01: $count"

