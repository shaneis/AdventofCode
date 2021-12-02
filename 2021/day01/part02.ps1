#! /usr/bin/pwsh

[CmdletBinding()]
param (
   $Path,

   $Groups = 3
)

$Measurements = Get-Content -Path $Path
$GroupLetter = 1

$IncreasedCounter = 0
for ($i = 3; $i -lt $Measurements.Count; $i++) {
   if ([int]$Measurements[$i] -gt [int]$Measurements[$i - 3]) {$IncreasedCounter++}
}
"Increased $IncreasedCounter times"