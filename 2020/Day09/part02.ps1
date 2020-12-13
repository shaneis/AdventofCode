function Get-ContiguousRange {
    [CmdletBinding()]
    param (
        [Parameter(ValueFromPipeline)]
        $Range,

        [Parameter()]
        $Search
    )
    
    begin {
        $newList = [Collections.Generic.List[Int64]]::new()
        $start = 0
    }
    
    process {
        foreach ($rNumber in $Range) {
            $newList.Add($rNumber) | Out-Null
        }
    }
    
    end {
        $sum = ($newList | Measure-Object -Sum).Sum
        $end = $newList.Count -1
        $sortedNewList = $newList | Sort-Object

        $result = [PSCustomObject]@{
            Search = $Search
            List = $newList
            ListSum = $sum
            Smallest = $sortedNewList[0]
            Largest = $sortedNewList[-1]
        }

        if ($result.Search -eq $result.ListSum) {
            $result
            return
        }

        if ($end -ne 1) {
            $newList = $newList[0..($end -1)]
            $newList | Get-ContiguousRange -Search $Search
        }
    }
}