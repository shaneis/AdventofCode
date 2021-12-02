#! /usr/bin/pwsh

[CmdletBinding()]
param (
    $Path
)

$PlannedCourse = Get-Content -Path $Path

$Horizontal = $Depth = 0

$Results = foreach ($Course in $PlannedCourse) {
    switch -Regex ($Course) {
        'forward' {
            $Horizontal += (-split $_)[1] -as [int]
        }
        'down' {
            $Depth += (-split $_)[1] -as [int]
        }
        'up' {
            $Depth += ((-split $_)[1] -as [int]) * -1
        }
    }

    [PSCustomObject] @{
        Horizontal = $Horizontal
        Depth = $Depth
    }
}

$Results

"Results: $($Results[-1].Horizontal) * $($Results[-1].Depth) = $($Results[-1].Horizontal * $Results[-1].Depth)"