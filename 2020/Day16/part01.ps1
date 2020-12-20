[CmdletBinding()]
param (
    [Parameter()]
    [string]
    $Path
)

$Content = (Get-Content -Path $Path -Raw) -split '(\r?\n){2}'

#region Rules
$Rules = $Content[0] -split '\r?\n'

$TargetNumbers = foreach ($r in $Rules) {

    $ParsedRule = [Regex]::Matches($r, '(?<rule_numbers>(\d+)-(\d+))')
    $vals = ($ParsedRule.Groups).Where({ $_.Name -eq 'rule_numbers'}).Value

    $vals | ForEach-Object -Process {
        $val = $_ -split ' '
        [PSCustomObject]@{
            First = ($val -split '-')[0] -as [int]
            Last = ($val -split '-')[1] -as [int]
        }
    }

}
#endregion

#region OurTicket

#endregion

#region OtherTickets
$Seen = [Collections.Generic.HashSet[int]]::new()

foreach ($ticket in ($Content[-1] -split '\r?\n')) {
    $Checked = [Collections.Generic.HashSet[int]]::new()
    $thisTicket = ($ticket -split ',').Where({ $_ -match '\d' })

    $thisTicket.ForEach({ 
        [int]$currentNumber = $_

        foreach ($target in $TargetNumbers) {
            if ($currentNumber -ge $target.First -and 
                $currentNumber -le $target.Last -and 
                -not $Seen.Contains($currentNumber)) {
                $Seen.Add($currentNumber) | Out-Null
                [PSCustomObject]@{
                    Number = $currentNumber
                    isValid = $true
                }
                break
            }
        }

        $TargetNumbers.ForEach({
            if (-not $Seen.Contains($currentNumber) -and -not $Checked.Contains($currentNumber)) {

                $Checked.Add($currentNumber) | Out-Null
                [PSCustomObject]@{
                    Number = $currentNumber
                    IsValid = $false
                }
            }
        })
    })
}
#endregion