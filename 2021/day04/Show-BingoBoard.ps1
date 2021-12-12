#! /usr/bin/pwsh
	
function Show-BingoBoard {
	[CmdletBinding()]
	param (
		[Parameter(
			Mandatory,
			ValueFromPipeline,
			ValueFromPipelineByPropertyName,
			Position = 0
		)]
		[Alias('Board')]
		[string] $BingoBoard,

		[Parameter(
			ValueFromPipeline,
			ValueFromPipelineByPropertyName,
			Position = 1
		)]
		[Alias('Numbers')]
		[int[]] $CalledNumbers
	)

	begin {
		$i = 0
		$Colour = 'White'

		Write-Verbose ($BingoBoard | Out-String)
	}

	process  {
		foreach ($Number in (-split ($BingoBoard -split '\r?\n'))) {
			Write-Debug $Number
			[bool] $nl = (($i + 1) % 5 -ne 0)	

			$Colour = if ($Number -in $CalledNumbers) {'Green'} else {'White'}

			$l = '{0:d2} ' -f ($Number -as [int])
			Write-Host $l -NoNewLine:$nl -ForegroundColor $Colour

			$i++		
		}	
	}
}

