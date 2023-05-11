package fuzz_switchboard_watcher

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"
    "github.com/cian911/switchboard/event"

    "github.com/cian911/switchboard/watcher"
)

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {
            
            case 0:
                var pathWatcher watcher.PathWatcher
                fuzzConsumer.GenerateStruct(&pathWatcher)
                interval, _ := fuzzConsumer.GetInt()
                if interval < 0 {
                    interval = -interval
                } else if interval == 0 {
                    interval = 1
                }

                testQ := watcher.NewQueue()
                var testEvent event.Event
                fuzzConsumer.GenerateStruct(&testEvent)
                repeat, _ := fuzzConsumer.GetInt()

                for i := 0; i < repeat; i++ {

                    testQ.Add(testEvent)
                }
                pathWatcher.Queue = testQ

                pathWatcher.Poll(interval)
                return 0

            case 1:
                testQ := watcher.NewQueue()
                var testEvent event.Event
                fuzzConsumer.GenerateStruct(&testEvent)
                repeat, _ := fuzzConsumer.GetInt()

                for i := 0; i < repeat; i++ {

                    testQ.Add(testEvent)
                }
                return 0

            case 2:
                testQ := watcher.NewQueue()
                var testEvent event.Event
                fuzzConsumer.GenerateStruct(&testEvent)
                repeat, _ := fuzzConsumer.GetInt()

                for i := 0; i < repeat; i++ {

                    testQ.Add(testEvent)
                }

                hash, _ := fuzzConsumer.GetString()
                testQ.Retrieve(hash)
                return 0

            case 3:
                testQ := watcher.NewQueue()
                var testEvent event.Event
                fuzzConsumer.GenerateStruct(&testEvent)
                repeat, _ := fuzzConsumer.GetInt()

                for i := 0; i < repeat; i++ {

                    testQ.Add(testEvent)
                }

                removingRepeat, _ := fuzzConsumer.GetInt()
                for i := 0; i < removingRepeat; i++ {
                    hash, _ := fuzzConsumer.GetString()
                    testQ.Remove(hash)
                }
                return 0

            case 4:
                testQ := watcher.NewQueue()
                var testEvent event.Event
                fuzzConsumer.GenerateStruct(&testEvent)
                repeat, _ := fuzzConsumer.GetInt()

                for i := 0; i < repeat; i++ {

                    testQ.Add(testEvent)
                }

                removingRepeat, _ := fuzzConsumer.GetInt()
                for i := 0; i < removingRepeat; i++ {
                    hash, _ := fuzzConsumer.GetString()
                    testQ.Remove(hash)
                }

                testQ.Size()
                return 0

            case 5:
                testQ := watcher.NewQueue()
                var testEvent event.Event
                fuzzConsumer.GenerateStruct(&testEvent)
                repeat, _ := fuzzConsumer.GetInt()

                for i := 0; i < repeat; i++ {

                    testQ.Add(testEvent)
                }

                removingRepeat, _ := fuzzConsumer.GetInt()
                for i := 0; i < removingRepeat; i++ {
                    hash, _ := fuzzConsumer.GetString()
                    testQ.Remove(hash)
                }

                testQ.Empty()
                return 0

            case 6:
                var testEvent event.Event
                fuzzConsumer.GenerateStruct(&testEvent)

                watcher.Hash(testEvent)
                return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}