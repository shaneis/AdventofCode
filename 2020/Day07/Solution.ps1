. $PSScriptRoot\part01.ps1
. $PSScriptRoot\part02.ps1

<# Part 01 #>
$part01Answer = Get-AOCBagHolders -Path $PSScriptRoot\input.txt -Colour 'shiny gold' |
    Select-Object -Property Bag -Unique |
    Measure-Object
'Part 01 Answer: {0}' -f $part01Answer.Count

<# Part 02 #><#
$part02Answer = Get-AOCGroupAnswersProper -Path $PSScriptRoot\input.txt |
    Measure-Object -Sum -Property AllYes
'Part 02 Answer: {0}' -f $part02Answer.Sum
#>
