[CmdletBinding()]
param (
    [Parameter()]
    [string]
    $Path
)

$Contents = (Get-Content -Path $path) -split '(\r?\n)' |
    Where-Object { -not [String]::IsNullOrWhiteSpace($_) }


$Width = ($Contents | Select-Object -First 1).Length
$Height = $Contents.Count

for ($i = 0; $i -lt $Height; $i++) {
    for ($j = 0; $j -lt $Width; $j++) {
        if ($Contents[$i][$j] -eq '.') { continue } 

        $above = ((($i + $Height) - 1) % $Height)
        $around = ((($j + $Width) - 1) % $Width)
        $occupiedSeats = 0
        $check = @(
            $Contents[$i][$j],
            ((($i + $Height) - 1) % $Height),
            ((($j + $Width) - 1) % $Width),
            $Contents[((($i + $Height) - 1 ) % $Height)][((($j + $Width) - 1) % $Width)]
        )
        'Check: {0} against height {1} and Width {2} : {3}' -f $check
    }
}