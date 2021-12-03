#! /usr/bin/pwsh

[CmdletBinding()]
param (
    $Path
)

$PlannedCourse = Get-Content -Path $Path

$Horizontal = $Depth = $Aim = 0

$Results = foreach ($Course in $PlannedCourse) {
    switch -Regex ($Course) {
        'forward' {
            $Horizontal += (-split $_)[1] -as [int]
            $Depth += $Aim * ((-split $_)[1] -as [int])
        }
        'down' {
            $Aim += (-split $_)[1] -as [int]
        }
        'up' {
            $Aim += ((-split $_)[1] -as [int]) * -1
        }
    }

    [PSCustomObject] @{
        Course = $Course
        Horizontal = $Horizontal
        Depth = $Depth
        Aim = $Aim
    }
}

$Results

"Results: $($Results[-1].Horizontal) * $($Results[-1].Depth) = $($Results[-1].Horizontal * $Results[-1].Depth)"
