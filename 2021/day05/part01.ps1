#! /usr/bin/pwsh

[CmdletBinding()]
param(
    $Path
)

$VentLines = Get-Content -Path $Path

$VentObjectID = 1
$VentLineObjects = foreach ($VentLine in $VentLines) {
    $FirstPart, $SecondPart = $VentLine -split '->'
    [int] $x1, [int] $y1 = $FirstPart -split ','
    [int] $x2, [int] $y2 = $SecondPart -split ',' 

    [PSCustomObject] @{
        VentObjectID = $VentObjectID++
        FullLine     = $VentLine
        x1           = $x1
        y1           = $y1
        x2           = $x2
        y2           = $y2
    }
}

foreach ($VentLineObject in $VentLineObjects) {
    [int] $startX = $VentLineObject.x1
    [int] $startY = $VentLineObject.y1
    [int] $endX = $VentLineObject.x2
    [int] $endY = $VentLineObject.y2

    do {
        [PSCustomObject]@{
            VentObjectID   = $VentLineObject.VentObjectID
            VentObjectLine = $VentLineObject.FullLine
            X              = $startX
            Y              = $startY
        }

        if ($startX -lt $endX) {
            $startX++
        }
        elseif ($startX -gt $endX) {
            $startX--
        }
        if ($startY -lt $endY) {
            $startY++
        }
        elseif ($startY -gt $endY) {
            $startY--
        }
    } while ($startX -ne $endX -or $startY -ne $endY)


    # Can't be bothered figuring out the off by one error
    if ($startX -ne $endX) { $startX -lt $endX ? ($startX++) : ($startX--) }
    if ($startY -ne $endY) { $startY -lt $endY ? ($startY++) : ($startY--) }

    [PSCustomObject]@{
        VentObjectID   = $VentLineObject.VentObjectID
        VentObjectLine = $VentLineObject.FullLine
        X              = $startX
        Y              = $startY
    }
}