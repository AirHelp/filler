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

@test "filler-run" {
    export TEST1="blabla"
    run $filler --dir test/output --ext tpl
    [ $status -eq 0 ]
    run grep "blabla" test/output/a.conf
    [ $status -eq 0 ]
    run grep "TEST2 is missing" test/output/b.conf
    [ $status -eq 0 ]
}

#@test "template-and-var-append" {
#  run $treasury template --src test/resources/bats-source.secret.tpl --dst test/output/bats-output.secret --append 'key1:treasury'
#  [ $status -eq 0 ]
#  run grep "key1=secret1treasury" test/output/bats-output.secret
#  [ $status -eq 0 ]
#}   


