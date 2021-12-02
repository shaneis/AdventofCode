#! /usr/bin/pwsh

[CmdletBinding()]
param (
   $Path,

   $Groups = 3
)

$Measurements = Get-Content -Path $Path

$IncreasedCounter = 0
for ($i = $Groups; $i -lt $Measurements.Count; $i++) {
   if ([int]$Measurements[$i] -gt [int]$Measurements[$i - $Groups]) {$IncreasedCounter++}
}
"Increased $IncreasedCounter times"