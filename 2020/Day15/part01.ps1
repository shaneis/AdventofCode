[CmdletBinding()]
param (
    [Parameter(ValueFromPipeline)]
    [string]
    $Path,

    [Parameter(ValueFromPipeline)]
    [long]
    $Target = 2020,

    [Parameter(ValueFromPipeline)]
    [switch]
    $Quiet
)

$InputContent = (Get-Content -Path $Path) -split ','

$Numbers = [Collections.Generic.List[psobject]]::new()
$Seen = [Collections.Generic.HashSet[int]]::new()
$ActualTarget = $Target - 1

foreach ($i in 0..$ActualTarget) {
    if (-not $Quiet) {
        Write-Progress -Activity DinnerTime -Status "$(($i / $ActualTarget) * 100)% complete" -PercentComplete (($i / $ActualTarget) * 100)
    }

    if ($i -lt $InputContent.Count) {
        $Numbers.Add(
            [PSCustomObject]@{
                Turn = $i + 1
                Number = $InputContent[$i]
            })
        if ($null -ne $lastNumberSpoken ) {
            $Seen.Add($lastNumberSpoken) | Out-Null
        }
        $lastNumberSpoken = $InputContent[$i]
        continue
    }
    
    if (-not $Seen.Contains($lastNumberSpoken)) {
        $Numbers.Add(
            [PSCustomObject]@{
                Turn = $i + 1
                Number = 0
            })
        $Seen.Add($lastNumberSpoken) | Out-Null
        $lastNumberSpoken = 0
    } else {
        $lastSeen = $Numbers |
            Where-Object Number -eq $lastNumberSpoken |
            Select-Object -Last 2

        $FirstLast = ($lastSeen[-1]).Turn
        $SecondLast = ($lastSeen[-2]).Turn

        $Numbers.Add(
            [PSCustomObject]@{
                Turn = $i + 1
                Number = $FirstLast - $SecondLast
            })
        $Seen.Add($lastNumberSpoken) | Out-Null
        $lastNumberSpoken = $FirstLast - $SecondLast
    }
}

$Numbers | Where-Object Turn -eq $Target