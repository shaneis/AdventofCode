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
    
    $char_string = [Text.StringBuilder]::new()
    for ($i = -1; $i -ge (0 - $Mask.Length); $i--) {
        Write-PSFMessage -Level Verbose -Message "Index: $i - Comparing $($ValueDecimal[$i]) against $($Mask[$i])"
        $char = switch ($Mask[$i]) {
            '0' { '0' }
            '1' { '1' }
            'X' { if ($null -ne $ValueDecimal[$i]) { $ValueDecimal[$i] } else { '0' }}
        }
        Write-PSFMessage -Level Verbose -Message "Adding $char to $($char_string.ToString())"
        $char_string.Insert(0, $char) | Out-Null
    }
    <#[PSCustomObject]@{
        InputValue = -join $ValueDecimal
        Mask = $Mask
        OutputValue = $char_string.ToString()
        ConvertedDecimal = [Convert]::ToInt32($char_String, 2)
        Location = $Location.Local
    }#>
    $Program[$Location.Local] = [Convert]::ToInt64($char_string, 2)
}

$Program.GetEnumerator() | Measure-Object -Sum -Property Value