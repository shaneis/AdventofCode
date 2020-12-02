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

    $Capture = $Capture | Select-Object -Property *,
        @{
            Name = 'JustLetter'
            Expression = { $PSItem.String -replace "[^$($PSItem.Letter)]" }
        },
        @{
            Name = 'Occurences'
            Expression = { ($PSItem.String -replace "[^$($PSItem.Letter)]").Length }
        }

    $Capture | Select-Object -Property *, 
        @{
            Name = 'Passes'

            Expression = { $PSItem.Occurences -ge [int]$PSItem.lower_bound -and $PSItem.Occurences -le [int]$PSItem.upper_bound }
        }
} 

Remove-Variable Samples, Capture, Sample
