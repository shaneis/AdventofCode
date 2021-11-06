[CmdletBinding()]
param (
	[string] $Path
)

#region import functions
. /home/soneill/git/AdventofCode/2019/day03/ConvertTo-WireInstructions.ps1
. /home/soneill/git/AdventofCode/2019/day03/ConvertTo-WirePath.ps1
#endregion

$TestLines = Get-Content -Path $Path

$Wire01, $Wire02 = $TestLines -split '\r?\n'

Write-Verbose "Wire 01: [$Wire01]"
Write-Verbose "Wire 02: [$Wire02]"

$Wire01Instructions = ConvertTo-WireInstructions -Wire $Wire01
$Wire02Instructions = ConvertTo-WireInstructions -Wire $Wire02

$Wire01Path = ConvertTo-WirePath -WireInstruction $Wire01Instructions
$Wire02Path = ConvertTo-WirePath -WireInstruction $Wire02Instructions

$Results = foreach ($Instr in $Wire01Path) {
	Write-Verbose "Going through: [$Instr]"
	foreach ($SecondPath in $Wire02Path) {
		if ($Instr.X -eq 0 -and $Instr.Y -eq 0) {continue}

		if ($Instr.X -eq $SecondPath.X -and $Instr.Y -eq $SecondPath.Y) {

			$ManhattanDistance = [Math]::Abs($SecondPath.X) + [Math]::Abs($SecondPath.Y)

			[PSCustomObject]@{
				X = $SecondPath.X
				Y = $SecondPath.Y
				ManhattanDistance = $ManhattanDistance
			}
		}
	}
}

$Results | Sort-Object -Property ManhattanDistance
