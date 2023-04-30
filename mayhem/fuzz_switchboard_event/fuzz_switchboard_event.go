package fuzz_switchboard_event

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/cian911/switchboard/event"
)

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {
            
            case 0:
                var testEvent event.Event
                fuzzConsumer.GenerateStruct(&testEvent)
                path, _ := fuzzConsumer.GetString()
                file, _ := fuzzConsumer.GetString()

                testEvent.Move(path, file)
                return 0

            case 1:
                var testEvent event.Event
                fuzzConsumer.GenerateStruct(&testEvent)
                ext, _ := fuzzConsumer.GetString()

                testEvent.IsValidEvent(ext)
                return 0

        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}