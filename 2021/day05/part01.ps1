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

$AllVentLines = foreach ($VentLineObject in $VentLineObjects) {
    [int] $startX = $VentLineObject.x1
    [int] $startY = $VentLineObject.y1
    [int] $EndX = $VentLineObject.x2
    [int] $EndY = $VentLineObject.y2
    [int] $OrigX = $VentLineObject.x1
    [int] $OrigY = $VentLineObject.y1

    do {
        [PSCustomObject]@{
            VentObjectID   = $VentLineObject.VentObjectID
            VentObjectLine = $VentLineObject.FullLine
            StartX         = $OrigX
            StartY         = $OrigY
            EndX           = $EndX
            EndY           = $EndY
            X              = $startX
            Y              = $startY
        }

        if ($startX -lt $EndX) {
            $startX++
        }
        elseif ($startX -gt $EndX) {
            $startX--
        }
        if ($startY -lt $EndY) {
            $startY++
        }
        elseif ($startY -gt $EndY) {
            $startY--
        }
    } while ($startX -ne $EndX -or $startY -ne $EndY)


    # Can't be bothered figuring out the off by one error
    if ($startX -ne $EndX) { $startX -lt $EndX ? ($startX++) : ($startX--) }
    if ($startY -ne $EndY) { $startY -lt $EndY ? ($startY++) : ($startY--) }

    [PSCustomObject]@{
        VentObjectID   = $VentLineObject.VentObjectID
        VentObjectLine = $VentLineObject.FullLine
        StartX         = $OrigX
        StartY         = $OrigY
        EndX           = $EndX
        EndY           = $EndY
        X              = $startX
        Y              = $startY
    }
}

$HVLines = foreach ($VL in $AllVentLines) {
    # Only interested in horizontal or vertical lines
    if ($VL.StartX -eq $VL.EndX -or $VL.StartY -eq $VL.EndY) {
        $VL
    }
}

($HVLines |
    Group-Object -Property X, Y -NoElement |
    Where-Object Count -gt 1).Count
