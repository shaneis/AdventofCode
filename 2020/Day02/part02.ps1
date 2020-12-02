[bool]$Test = $false

[string[]]$Samples = if ($Test) {
    '1 - 3 a: abcde'
    '1 - 3 b: cdefg'
    '2 - 9 c: ccccccccc'
} else {
    Get-Content -Path "$PSScriptRoot\day02_part01.txt"
}

foreach ($sample in $Samples) {
    $parseLine = ('(?<Amount>(?<lower_bound>\d*)\s?-\s?(?<upper_bound>\d*))',
        '(?<Letter>[a-z]):',
        '(?<String>\w+)') -join '\s'

    $sample -match $parseLine | Out-Null
    $Matches.Remove(0)
    $Capture = [PSCustomObject]$Matches

    $firstPosition = ($Capture.lower_bound -as [int]) - 1
    $lastPosition = ($Capture.upper_bound -as [int]) - 1

    [bool]$isValid = $false
    if ($Capture.String[$firstPosition] -eq $Capture.Letter -or
        $Capture.String[$lastPosition] -eq $Capture.Letter) {
            $isValid = $true
        }
    if ($Capture.String[$firstPosition] -eq $Capture.String[$lastPosition]) {
        $isValid = $false
    }

    $Capture | Select-Object -Property *, @{ Name = 'IsValid'; Expression = { $isValid }}
} 

Remove-Variable Samples, Capture, Sample
