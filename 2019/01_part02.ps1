function Get-FuelPerModuleMassWithFuel {
    [CmdletBinding()]
    param (
        [Parameter(ValueFromPipeline)]
        [int[]]$Mass
    )

    begin {
        
        $GetFuelModuleMassFunction = Join-Path -Path $PSScriptRoot -ChildPath 01_part01.ps1
        . $GetFuelModuleMassFunction
    }

    process {
        
        foreach ($moduleMass in $Mass) {
            $ResultStack = [System.Collections.Generic.Stack[int]]::new()
            
            $ResultStack.Push((Get-FuelPerModuleMass -Mass $Mass).Fuel)

            while ($ResultStack.Peek() -gt 0) {
                $ResultStack.Push((Get-FuelPerModuleMass -Mass $ResultStack.Peek()).Fuel)
            }

            [PSCustomObject]@{
                InitialMass = $Mass
                Fuel = $ResultStack
                TotalFuel = ($ResultStack | Measure-Object -Sum).Sum
            }
        }   
    }
}
