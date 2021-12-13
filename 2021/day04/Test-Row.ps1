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
			$BoardNumbers = $BoardRow -split ' ' | Where-Object {$_ -as [int] -or $_ -eq '0'}
			foreach ($n in $BoardNumbers) {
				Write-Verbose "Row: $Row - Comparing: $n_int against $($CalledNumbers -join ',')"
					[int] $n_int = $n
					 if ($n_int -notin $CalledNumbers) {
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
