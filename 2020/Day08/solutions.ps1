[CmdletBinding()]
param ($path)

. .\part02.ps1

$con = (Get-Content -Path $path -Raw) -split '(\r?\n)' |
            Where-Object { -not [string]::IsNullOrWhiteSpace($_) }

foreach ($c in 0..($con.Count -1)) {

    $con_clone = $con.Clone()

    if ($con_clone[$c] -match 'jmp') {
        Write-PSFMessage -Level Host -Message "Chaging line $c : jmp-->nop: $($con_clone[$c])"
        $con_clone[$c] = $con_clone[$c] -replace 'jmp', 'nop'
        $tmpFile = New-TemporaryFile
        $con_clone | Set-Content -path $tmpFile

        #Get-Content $tmpFile

        $ws = Get-AOCAccValue_Fix -Path $tmpFile
        if ($ws | Where-Object WasSuccessful) {
            Write-PSFMessage -Level Host -Message "$c - jmp - $($ws.WasSuccessful)"
            $ws | Format-Table
            return
        }
    }

    $con_clone = $con.Clone()
    if ($con_clone[$c] -match 'nop') {
        Write-PSFMessage -Level Host -Message "Changing line $c : npp-->jmp: $($con_clone[$c])"
        $con_clone[$c] = $con_clone[$c] -replace 'nop', 'jmp'
        $tmpFile = New-TemporaryFile
        $con_clone | Out-File -FilePath $tmpFile -Force

        $ws = Get-AOCAccValue_Fix -Path $tmpFile
        if ($ws | Where-Object WasSuccessful) {
            Write-PSFMessage -Level Host -Message "$c - nop - $($ws.WasSuccessful)"
            $ws | Format-Table
            return
        }
    }

}

