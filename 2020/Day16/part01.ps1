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
Write-Verbose -Message "$($TargetNumbers | Out-String)"
#endregion

#region OurTicket

#endregion

#region OtherTickets
$Tickets = $Content[-1] -split '\r?\n'

$Seen = [Collections.Generic.HashSet[int]]::new()

foreach ($ticket in $Tickets) {

    $thisTicket = ($Tickets -split ',').Where({ $_ -match '\d' })

    $thisTicket.ForEach({ 
        [int]$currentNumber = $_

        foreach ($target in $TargetNumbers) {
            if ($currentNumber -ge $target.First -and $currentNumber -le $target.Last) {
                $Seen.Add($currentNumber) | Out-Null
            }
        }
    })

}

foreach ($oth in ($Tickets -split ',').Where({ $_ -match '\d' })) {
    Write-Verbose -Message "$oth"
    if (-not $Seen.Contains($oth)) {
        [PSCustomObject]@{
            Number = $oth
            IsValid = $false
        }
    }
}


#endregion