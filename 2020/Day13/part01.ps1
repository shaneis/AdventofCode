[CmdletBinding()]
param (
    [Parameter()]
    [string]
    $Path
)

$InputContent = Get-Content -Path $Path

$OurTimestamp = $InputContent[0]
$Buses = $InputContent[1] -split ',' | Where-Object { $_ -match '[^x\s]'}

$Results = foreach ($Bus in $Buses) {
    #'Going to multiply bus {0} by {1} for timestamp {2}' -f $Bus, ([Math]::Ceiling($OurTimestamp / $Bus)), $OurTimestamp
    $NextOccurence = [Math]::Ceiling($OurTimestamp / $Bus)
    [PSCustomObject]@{
        OurTimestamp = $OurTimestamp
        Bus = $Bus
        NextArrival = [int]$Bus * $NextOccurence
    }
}

$EarliestBusTime = $Results | Measure-Object -Property NextArrival -Minimum

$EarlistBus = $Results | Where-Object NextArrival -eq $EarliestBusTime.Minimum

$DifferenceInMinutes = $EarliestBusTime.Minimum - $OurTimestamp
'Bus {0} * DifferenceInMinutes {1} = {2}' -f $EarlistBus.Bus, $DifferenceInMinutes, ($DifferenceInMinutes * $EarlistBus.Bus)