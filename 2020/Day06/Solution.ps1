. $PSScriptRoot\part01.ps1

<# Part 01 #>
$part01Answer = Get-AOCGroupAnswers -Path $PSScriptRoot\input.txt |
    Measure-Object -Sum -Property AllYes
'Part 01 Answer: {0}' -f $part01Answer.Sum