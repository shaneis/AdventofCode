BeforeAll -ScriptBlock {
    . "$PSScriptRoot\part01.ps1"
}

Describe -Name 'part01' -Fixture {
    Context -Name 'Given example' -Fixture {
        It -Name 'returns row <SeatRow>, column <SeatColumn>, and ID <SeatID> for <SeatCode>' -TestCases @(
            @{ SeatCode = 'FBFBBFFRLR'; SeatRow = 44; SeatColumn = 5; SeatID = 357 }
            @{ SeatCode = 'BFFFBBFRRR'; SeatRow = 70; SeatColumn = 7; SeatID = 567 }
            @{ SeatCode = 'FFFBBBFRRR'; SeatRow = 14; SeatColumn = 7; SeatID = 119 }
            @{ SeatCode = 'BBFFBBFRLL'; SeatRow = 102; SeatColumn = 4; SeatID = 820 }
        ) -Test {
            param (
                $SeatCode,
                $SeatRow,
                $SeatColumn,
                $SeatID
            )
            $seatDetails = Get-SeatNumber -SeatCode $SeatCode
            $seatDetails.SeatRow | Should -Be $SeatRow
            $seatDetails.SeatColumn | Should -Be $SeatColumn
            $seatDetails.SeatID | Should -Be $SeatID
        }
    }
}
