. $PSScriptRoot\part01.ps1

<# Part 01 #>
$AllSeats = foreach ($Seat in Get-Content -Path "$PSScriptRoot\input_part01.txt") {
    Get-SeatNumber -SeatCode $Seat
}
$highestSeatID = $AllSeats | Measure-Object -Maximum -Property SeatID
"Highest Seat ID: $($highestSeatID.Maximum)"

<# Part 02 #>
$permutations = 0..127 | ForEach-Object -Process {
    $row = $PSItem
    0..7 | ForEach-Object -Process { 
        [PSCustomObject]@{
            SeatRow = $row
            SeatColumn = $PSItem
            SeatID = ($row * 8) + $PSItem
        }
    }
}

$MySeat = $permutations |
    Where-Object SeatRow -notin '0','127' |
    Where-Object SeatID -notin $AllSeats.SeatID |
    Where-Object { 
        ($PSItem.SeatID + 1) -in $AllSeats.SeatID -and
        ($PSItem.SeatID - 1) -in $AllSeats.SeatID
    }
"My Seat ID: $($MySeat.SeatID)"