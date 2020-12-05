function Get-SeatNumber {
    [CmdletBinding()]
    param (
        [Parameter(Mandatory)]
        [string]$SeatCode
    )
    
    begin {
        $rowArray = 0..127
        $columnArray = 0..7
    }
    
    process {
        
        $workingArray = $rowArray.Clone()
        foreach ($char in $SeatCode.ToCharArray() | Select-Object -First 7) {

            [int]$halfWay = ($workingArray.Count / 2) 

            switch ($char) {
                'F' { $start = 0; $end = $halfway - 1 }
                'B' { $start = $halfway; $end = $workingArray.Count }
                default { Write-Error 'Not F or B'}
            }

            $workingArray = $workingArray[$start..$end]
        }

        $seatArray = $columnArray.Clone()
        foreach ($seatChar in $SeatCode.ToCharArray() | Select-Object -Last 3) {

            [int]$halfway = ($seatArray.Count / 2)

            switch ($seatChar) {
                'L' { $start = 0; $end = $halfway - 1 }
                'R' { $start = $halfway; $end = $seatArray.Count }
            }

            $seatArray = $seatArray[$start..$end]
        }

        [PSCustomObject]@{
            SeatCode = $SeatCode
            SeatRow = $workingArray[0]
            SeatColumn = $seatArray[0]
            SeatID = ((1 * $workingArray[0]) * 8) + $seatArray[0]
        }
    }
    
    end {
        
    }
}