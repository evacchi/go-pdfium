;;repro.wat

  (func (;5011;) (type 2) (param i32 i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get 0
    i32.const -64
    i32.add
    local.tee 9
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        i32.load
        i32.const 1
        i32.ne
        br_if 0 (;@2;)
        local.get 0
        i32.const 20
        i32.add
        local.set 13
        loop  ;; label = @3
          local.get 0
          i32.load offset=40
          local.tee 3
          if (result i32)  ;; label = @4
            local.get 3
          else
            local.get 0
            i32.load offset=36
            local.tee 3
            local.get 0
            i32.load offset=4
            local.tee 4
            i32.load offset=16
            local.get 4
            i32.load offset=12
            local.tee 2
            i32.sub
            i32.const 28
            i32.div_s
            i32.ge_u
            if  ;; label = @5
              local.get 0
              i32.const 2
              i32.store
              br 3 (;@2;)
            end
            local.get 0
            local.get 2
            local.get 3
            i32.const 28
            i32.mul
            i32.add
            local.tee 3
            i32.store offset=40
            local.get 3
            i32.load
            local.tee 3
            i32.load offset=84
            local.tee 11
            local.get 3
            i32.load offset=96
            local.get 3
            i32.load offset=100
            i32.add
            local.tee 8
            i32.const 8
            i32.shr_u
            i32.const 16777212
            i32.and
            i32.add
            local.set 2
            local.get 0
            local.get 11
            local.get 3
            i32.load offset=88
            i32.eq
            if (result i32)  ;; label = @5
              i32.const 0
            else
              local.get 2
              i32.load
              local.get 8
              i32.const 1023
              i32.and
              i32.const 2
              i32.shl
              i32.add
            end
            i32.store offset=48
            local.get 0
            local.get 2
            i32.store offset=44
            i32.const 144
            call 106
            local.get 4
            local.get 0
            i32.load offset=8
            call 4134
            local.set 4
            local.get 0
            i32.load offset=16
            local.set 3
            local.get 0
            local.get 4
            i32.store offset=16
            local.get 3
            if  ;; label = @5
              local.get 3
              call 4135
              call 76
            end
            block  ;; label = @5
              local.get 0
              i32.load offset=12
              local.tee 3
              i32.eqz
              br_if 0 (;@5;)
              local.get 0
              i32.load offset=16
              local.tee 4
              local.get 3
              i32.load16_u
              i32.store16
              local.get 4
              i32.const 2
              i32.add
              local.get 3
              i32.const 2
              i32.add
              call 4087
              local.get 4
              local.get 3
              i64.load offset=20 align=4
              i64.store offset=20 align=4
              local.get 4
              local.get 3
              i64.load offset=12 align=4
              i64.store offset=12 align=4
              local.get 4
              i32.load offset=28
              local.tee 2
              local.get 3
              i32.load offset=28
              local.tee 3
              i32.eq
              br_if 0 (;@5;)
              local.get 3
              if  ;; label = @6
                local.get 3
                local.get 3
                i32.load offset=4
                i32.const 1
                i32.add
                local.tee 2
                i32.store offset=4
                local.get 2
                i32.eqz
                br_if 5 (;@1;)
                local.get 4
                i32.load offset=28
                local.set 2
              end
              local.get 4
              local.get 3
              i32.store offset=28
              local.get 2
              i32.eqz
              br_if 0 (;@5;)
              local.get 2
              i32.load offset=4
              local.tee 3
              i32.eqz
              br_if 4 (;@1;)
              local.get 2
              local.get 3
              i32.const 1
              i32.sub
              local.tee 3
              i32.store offset=4
              local.get 3
              br_if 0 (;@5;)
              local.get 2
              local.get 2
              i32.load
              i32.load offset=4
              call_indirect (type 0)
            end
            local.get 0
            i32.load offset=16
            i32.const 124
            i32.add
            local.get 0
            i32.load offset=40
            i32.load
            i32.const 52
            i32.add
            call 3131
            drop
            local.get 0
            i32.load offset=16
            i32.const 0
            i32.const 0
            call 4136
            local.get 0
            i32.load offset=8
            call 2117
            local.get 9
            i32.const 24
            i32.add
            local.get 0
            i32.load offset=40
            i32.const 4
            i32.add
            call 2026
            local.get 9
            i32.const 48
            i32.add
            local.get 9
            i32.const 24
            i32.add
            local.get 9
            i32.const 8
            i32.add
            local.get 0
            i32.load offset=8
            i32.const 28
            i32.add
            call 2009
            call 2032
            local.get 13
            local.get 9
            i64.load offset=56 align=4
            i64.store offset=8 align=4
            local.get 13
            local.get 9
            i64.load offset=48 align=4
            i64.store align=4
            local.get 0
            i32.load offset=40
          end
          i32.load
          local.tee 3
          i32.load offset=96
          local.set 2
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 3
                i32.load offset=84
                local.tee 4
                local.get 3
                i32.load offset=88
                i32.eq
                if  ;; label = @7
                  i32.const 0
                  local.set 11
                  local.get 0
                  i32.load offset=48
                  local.tee 8
                  br_if 1 (;@6;)
                  local.get 4
                  local.get 2
                  i32.const 8
                  i32.shr_u
                  i32.const 16777212
                  i32.and
                  i32.add
                  local.set 4
                  i32.const 0
                  local.set 3
                  br 3 (;@4;)
                end
                local.get 0
                i32.load offset=48
                local.tee 8
                local.get 4
                local.get 3
                i32.load offset=100
                local.get 2
                i32.add
                local.tee 3
                i32.const 8
                i32.shr_u
                i32.const 16777212
                i32.and
                i32.add
                i32.load
                local.get 3
                i32.const 1023
                i32.and
                i32.const 2
                i32.shl
                i32.add
                local.tee 11
                i32.eq
                br_if 1 (;@5;)
              end
              local.get 8
              i32.const 4
              i32.add
              local.tee 3
              local.get 0
              i32.load offset=44
              local.tee 2
              i32.load
              i32.sub
              i32.const 4096
              i32.ne
              if  ;; label = @6
                local.get 2
                local.set 4
                br 2 (;@4;)
              end
              local.get 2
              i32.const 4
              i32.add
              local.set 4
              local.get 2
              i32.load offset=4
              local.set 3
              br 1 (;@4;)
            end
            local.get 4
            local.get 2
            i32.const 8
            i32.shr_u
            i32.const 16777212
            i32.and
            i32.add
            local.tee 4
            i32.load
            local.get 2
            i32.const 1023
            i32.and
            i32.const 2
            i32.shl
            i32.add
            local.set 3
          end
          i32.const 100
          local.set 12
          i32.const 0
          local.set 8
          loop  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 3
                local.get 11
                i32.ne
                if  ;; label = @7
                  block  ;; label = @8
                    local.get 3
                    i32.load
                    local.tee 2
                    i32.eqz
                    br_if 0 (;@8;)
                    local.get 2
                    f32.load offset=24
                    local.get 0
                    f32.load offset=28
                    f32.le
                    i32.eqz
                    br_if 0 (;@8;)
                    local.get 2
                    f32.load offset=32
                    local.get 13
                    f32.load
                    f32.ge
                    i32.eqz
                    br_if 0 (;@8;)
                    local.get 2
                    f32.load offset=28
                    local.get 0
                    f32.load offset=32
                    f32.le
                    i32.eqz
                    br_if 0 (;@8;)
                    local.get 2
                    f32.load offset=36
                    local.get 0
                    f32.load offset=24
                    f32.ge
                    i32.eqz
                    br_if 0 (;@8;)
                    local.get 0
                    i32.load offset=12
                    i32.load8_u offset=6
                    i32.eqz
                    br_if 2 (;@6;)
                    local.get 2
                    local.get 2
                    i32.load
                    i32.load offset=24
                    call_indirect (type 1)
                    i32.eqz
                    br_if 2 (;@6;)
                    local.get 2
                    local.get 2
                    i32.load
                    i32.load offset=52
                    call_indirect (type 1)
                    call 3754
                    local.tee 5
                    i32.load offset=4
                    local.tee 6
                    i32.eqz
                    br_if 7 (;@1;)
                    local.get 5
                    i32.load8_u offset=21
                    local.set 7
                    local.get 5
                    local.get 6
                    i32.const 1
                    i32.sub
                    local.tee 6
                    i32.store offset=4
                    block  ;; label = @9
                      local.get 6
                      i32.eqz
                      if  ;; label = @10
                        local.get 5
                        local.get 5
                        i32.load
                        i32.load offset=4
                        call_indirect (type 0)
                        local.get 7
                        br_if 1 (;@9;)
                        br 4 (;@6;)
                      end
                      local.get 7
                      i32.eqz
                      br_if 3 (;@6;)
                    end
                    i32.const 1
                    local.set 8
                    local.get 0
                    i32.load offset=8
                    i32.load8_u offset=24
                    i32.eqz
                    br_if 2 (;@6;)
                    local.get 0
                    local.get 3
                    i32.store offset=48
                    local.get 0
                    local.get 4
                    i32.store offset=44
                    local.get 0
                    i32.load offset=16
                    local.get 2
                    i32.const 4
                    i32.add
                    local.get 0
                    i32.load offset=40
                    i32.const 4
                    i32.add
                    call 4138
                    br 6 (;@2;)
                  end
                  local.get 0
                  local.get 3
                  i32.store offset=48
                  local.get 0
                  local.get 4
                  i32.store offset=44
                  br 2 (;@5;)
                end
                block  ;; label = @7
                  local.get 0
                  i32.load offset=40
                  i32.load
                  local.tee 3
                  i32.load8_u offset=55
                  i32.const 2
                  i32.eq
                  if  ;; label = @8
                    local.get 0
                    i32.load offset=16
                    local.set 3
                    local.get 0
                    i32.const 0
                    i32.store offset=16
                    local.get 3
                    if  ;; label = @9
                      local.get 3
                      call 4135
                      call 76
                    end
                    local.get 0
                    i32.load offset=8
                    i32.const 0
                    call 2114
                    local.get 0
                    i32.const 0
                    i32.store offset=40
                    local.get 0
                    local.get 0
                    i32.load offset=36
                    i32.const 1
                    i32.add
                    i32.store offset=36
                    local.get 8
                    i32.const 1
                    i32.and
                    br_if 6 (;@2;)
                    local.get 1
                    i32.eqz
                    br_if 1 (;@7;)
                    local.get 1
                    local.get 1
                    i32.load
                    i32.load offset=8
                    call_indirect (type 1)
                    i32.eqz
                    br_if 1 (;@7;)
                    br 6 (;@2;)
                  end
                  local.get 8
                  i32.const 1
                  i32.and
                  br_if 5 (;@2;)
                  local.get 3
                  local.get 1
                  call 3884
                  local.get 0
                  i32.load offset=40
                  i32.load
                  i32.load8_u offset=55
                  i32.const 2
                  i32.ne
                  br_if 5 (;@2;)
                end
                local.get 0
                i32.load
                i32.const 1
                i32.eq
                br_if 3 (;@3;)
                br 4 (;@2;)
              end
              block (result i32)  ;; label = @6
                local.get 0
                i32.load offset=16
                local.set 5
                local.get 0
                i32.load offset=40
                i32.const 4
                i32.add
                local.set 10
                local.get 2
                i32.const 4
                i32.add
                local.set 14
                block  ;; label = @7
                  block  ;; label = @8
                    loop  ;; label = @9
                      local.get 5
                      i32.load offset=116
                      local.tee 6
                      if  ;; label = @10
                        i32.const 1
                        local.set 7
                        local.get 6
                        local.get 1
                        call 4108
                        br_if 3 (;@7;)
                        block  ;; label = @11
                          local.get 5
                          i32.load offset=116
                          local.tee 6
                          i32.load8_u offset=98
                          if  ;; label = @12
                            local.get 5
                            i32.const 0
                            i32.store offset=116
                            br 1 (;@11;)
                          end
                          local.get 5
                          local.get 2
                          local.get 10
                          call 4144
                          local.get 5
                          i32.load offset=116
                          local.set 6
                          i32.const 0
                          local.set 7
                          local.get 5
                          i32.const 0
                          i32.store offset=116
                          local.get 6
                          i32.eqz
                          br_if 4 (;@7;)
                        end
                        local.get 6
                        call 4102
                        call 76
                        br 2 (;@8;)
                      end
                      local.get 5
                      local.get 2
                      i32.store offset=88
                      i32.const 0
                      local.set 7
                      local.get 5
                      local.get 2
                      call 4091
                      i32.eqz
                      br_if 2 (;@7;)
                      local.get 5
                      local.get 14
                      local.get 10
                      call 4138
                      local.get 5
                      local.get 2
                      local.get 10
                      call 4139
                      br_if 2 (;@7;)
                      local.get 2
                      local.get 2
                      i32.load
                      i32.load offset=24
                      call_indirect (type 1)
                      i32.eqz
                      if  ;; label = @10
                        local.get 5
                        local.get 2
                        local.get 10
                        call 4140
                        i32.const 0
                        br 4 (;@6;)
                      end
                      i32.const 100
                      call 106
                      local.get 5
                      call 4101
                      local.set 7
                      local.get 5
                      i32.load offset=116
                      local.set 6
                      local.get 5
                      local.get 7
                      i32.store offset=116
                      local.get 6
                      if (result i32)  ;; label = @10
                        local.get 6
                        call 4102
                        call 76
                        local.get 5
                        i32.load offset=116
                      else
                        local.get 7
                      end
                      local.get 2
                      local.get 2
                      i32.load
                      i32.load offset=52
                      call_indirect (type 1)
                      local.get 10
                      i32.const 0
                      i32.const 0
                      call 4107
                      br_if 0 (;@9;)
                    end
                    block  ;; label = @9
                      local.get 5
                      i32.load offset=116
                      local.tee 6
                      i32.load8_u offset=98
                      if  ;; label = @10
                        local.get 5
                        i32.const 0
                        i32.store offset=116
                        br 1 (;@9;)
                      end
                      local.get 5
                      local.get 2
                      local.get 10
                      call 4144
                      local.get 5
                      i32.load offset=116
                      local.set 6
                      i32.const 0
                      local.set 7
                      local.get 5
                      i32.const 0
                      i32.store offset=116
                      local.get 6
                      i32.eqz
                      br_if 2 (;@7;)
                    end
                    local.get 6
                    call 4102
                    call 76
                  end
                  i32.const 0
                  local.set 7
                end
                local.get 7
              end
              br_if 3 (;@2;)
              block  ;; label = @6
                local.get 2
                local.get 2
                i32.load
                i32.load offset=24
                call_indirect (type 1)
                i32.eqz
                br_if 0 (;@6;)
                local.get 0
                i32.load offset=16
                i32.load8_u offset=10
                i32.eqz
                br_if 0 (;@6;)
                local.get 0
                i32.load offset=4
                i32.load offset=8
                i32.const 104857600
                call 3559
              end
              block  ;; label = @6
                block  ;; label = @7
                  local.get 2
                  local.get 2
                  i32.load
                  i32.load offset=32
                  call_indirect (type 1)
                  i32.eqz
                  if  ;; label = @8
                    local.get 2
                    local.get 2
                    i32.load
                    i32.load offset=28
                    call_indirect (type 1)
                    i32.eqz
                    br_if 1 (;@7;)
                  end
                  local.get 0
                  local.get 3
                  i32.store offset=48
                  local.get 0
                  local.get 4
                  i32.store offset=44
                  local.get 1
                  br_if 1 (;@6;)
                  i32.const 100
                  local.set 12
                  br 2 (;@5;)
                end
                local.get 0
                local.get 3
                i32.store offset=48
                local.get 0
                local.get 4
                i32.store offset=44
                local.get 12
                i32.const 1
                i32.sub
                local.tee 2
                i32.const 100
                local.get 2
                select
                local.set 12
                local.get 2
                br_if 1 (;@5;)
                local.get 1
                i32.eqz
                br_if 1 (;@5;)
              end
              i32.const 100
              local.set 12
              local.get 1
              local.get 1
              i32.load
              i32.load offset=8
              call_indirect (type 1)
              br_if 3 (;@2;)
            end
            local.get 3
            i32.const 4
            i32.add
            local.tee 3
            local.get 4
            i32.load
            i32.sub
            i32.const 4096
            i32.eq
            if  ;; label = @5
              local.get 4
              i32.load offset=4
              local.set 3
              local.get 4
              i32.const 4
              i32.add
              local.set 4
            end
            local.get 8
            local.get 3
            local.get 11
            i32.ne
            i32.and
            i32.eqz
            br_if 0 (;@4;)
          end
        end
      end
      local.get 9
      i32.const -64
      i32.sub
      global.set 0
      return
    end
    unreachable)
