[CmdletBinding()]
param (
    [Parameter()]
    [string]
    $Path
)

$InputContent = Get-Content -Path $Path

<#
    $Memory, $Value = $InputContent -split ' = '
    $Memory
    $Value
#>
$Program = @{}

foreach ($Instruction in $InputContent) {
    $Memory, $Value = $Instruction -split ' = '

    if ($Memory -eq 'mask') {
        Write-PSFMessage -Level Host -Message ('Mask: {0}' -f $Value)
        $Mask = $value
        continue
    }

    $Location = $Memory -match ('mem\[(?<local>\d+)\]') | Select-Object -Property @{ Name = 'Local'; Expression = { $Matches['local'] }}

    $ValueDecimal = [Convert]::ToString($Value, 2).ToCharArray()
    
    
    for ($i = - 1; $i -ge (0 - $ValueDecimal.Count); $i--) {
        Write-PSFMessage -Level Host -Message "Index: $i - Comparing $($ValueDecimal[$i]) against $($Mask[$i])"
        
    }
}

#$Program