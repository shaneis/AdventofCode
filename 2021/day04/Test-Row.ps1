#! /usr/bin/pwsh

function Test-Row {
    [CmdletBinding()]
	param (
		[Parameter(
			Mandatory,
			ValueFromPipeline,
			ValueFromPipelineByPropertyName
		)]
		[Alias('BingoBoard')]
		[string] $Board,

		[Parameter(
			Mandatory,
			ValueFromPipelineByPropertyName
		)]
		[int] $Row,

		[Parameter(
			Mandatory,
			ValueFromPipelineByPropertyName
		)]
		[Alias('Numbers')]
		[int[]] $CalledNumbers
	)

	begin {
		$IsComplete = $true
	}

	process {
		$BoardRows = $Board -split '\r?\n' |
			Select-Object -First 1 -Skip ($Row -1)

		foreach ($BoardRow in $BoardRows) {
			foreach ($n in ($BoardRow -split ' ')) {
				if ($n -notin $CalledNumbers) {
					$IsComplete = $false
					break
				}
			}

			[PSCustomObject] @{
				BoardRow = $BoardRow
				Row = $Row
				Numbers = $CalledNumbers -join ','
				IsComplete = $IsComplete
			}
		}
	}	
}
