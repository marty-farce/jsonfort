        subroutine fort(input)
            implicit none
            type :: Fix
                INTEGER :: ADEP_1
                INTEGER :: ADEP_2
                INTEGER :: ADEP_3
                INTEGER :: ADEP_4
                INTEGER :: ADEP_5
            end type Fix
            type(Fix) :: input
            WRITE(*,*) input
            RETURN
        end