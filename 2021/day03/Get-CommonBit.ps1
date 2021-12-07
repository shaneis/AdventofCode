function Get-CommonBit {
    [CmdletBinding()]
    param (
        [Parameter(
            Mandatory,
            ValueFromPipeline,
            ValueFromPipelineByPropertyName
        )]
        $Data,

        $Position,

        [ValidateSet('O2', 'CO2')]
        [string] $Type
    )
    
    begin {
        $DataCollection = [Collections.Generic.List[string]]::new()
    }
    
    process {
        foreach ($Datum in $Data) {
            $null = $DataCollection.Add($Datum)
        }
    }
    
    end {
        $GroupedElements = $DataCollection | ForEach-Object -Process {
            $_[$Position]
        } |
        Group-Object -NoElement 

        if ($GroupedElements[0].Count -eq $GroupedElements[1].Count) {
            if ($Type -eq 'O2') {
                $GroupedElements | Where-Object Name -eq '1'
            }
            else {
                $GroupedElements | Where-Object Name -eq '0'
            }
        }
        else {
            if ($Type -eq 'O2') {
                $GroupedElements |
                Sort-Object Count -Descending |
                Select-Object -First 1
            }
            else {
                $GroupedElements |
                Sort-Object Count -Descending |
                Select-Object -Last 1
            }
        }
    }
}