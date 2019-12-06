function Get-FuelPerModuleMass {
    <#
    .SYNOPSIS
        Gets required fuel per the module mass.
    .DESCRIPTION
        Fuel required to launch a given module is based on its mass.
            Specifically, to find the fuel required for a module, take its mass, divide by three,
            round down, and subtract 2.
    .EXAMPLE
        PS C:\> Get-FuelPerModuleMass -Mass 12

        Mass Fuel
        ---- ----
          12    2
    .EXAMPLE
        PS C:\> Get-FuelPerModuleMass -Mass 12, 14

        Mass Fuel
        ---- ---- 
          12    2
          14    2
    .INPUTS
        [int[]]
    .OUTPUTS
        [PSCustomObject]
    #>
    [CmdletBinding()]
    param (
        [Parameter(Mandatory,
                   ValueFromPipeline)]
        [int[]]$Mass
    )

    process {
        foreach ($moduleMass in $Mass) {
            [PSCustomObject]@{
                Mass = $moduleMass
                Fuel = [Math]::Floor(($moduleMass / 3)) - 2
            }
        }
    }
}
