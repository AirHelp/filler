#!/usr/bin/env bats

filler=$PWD/filler

@test "Check that the filler binary is available" {
  run stat $filler
  [ $status -eq 0 ]
}

@test "help" {
  run $filler --help
  [ $status -eq 0 ]
}

@test "version" {
    run $filler version
    [ $status -eq 0 ]
    [[ ${lines[0]} =~ "Version: " ]]
}

@test "filler-run-without-ext" {
    export TEST1="blabla"
    run $filler --src test/output
    [ $status -eq 0 ]
    run grep "blabla" test/output/a.conf
    [ $status -eq 0 ]
    run grep "TEST2 is missing" test/output/b.conf
    [ $status -eq 0 ]
}

@test "filler-run-with-ext" {
    export TEST2="blabla"
    run $filler --src test/output --ext tpl_new
    [ $status -eq 0 ]
    run grep "blabla" test/output/c.conf
    [ $status -eq 0 ]
}

@test "filler-run-with-single-file" {
    export TEST1="blabla"
    run $filler --src test/output/d.conf.tpl_single --ext tpl_single
    [ $status -eq 0 ]
    run grep "blabla" test/output/d.conf
    [ $status -eq 0 ]
}



