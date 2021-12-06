function Get-MostCommonBit {
    [CmdletBinding()]
    param (
        [Parameter(
            Mandatory,
            ValueFromPipeline,
            ValueFromPipelineByPropertyName
        )]
        $Data,

        $Position
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
            $GroupedElements | Where-Object Name -eq '1'
        }
        else {
            $GroupedElements |
            Sort-Object Count -Descending |
            Select-Object -First 1
        }
    }
}