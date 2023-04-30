package fuzz_switchboard_utils

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/cian911/switchboard/utils"
)

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {
            
            case 0:
                path, _ := fuzzConsumer.GetString()

                utils.ExtractFileExt(path)
                return 0

            case 1:
                path, _ := fuzzConsumer.GetString()

                utils.ExtractPathWithoutExt(path)
                return 0

            case 2:
                path1, _ := fuzzConsumer.GetString()
                path2, _ := fuzzConsumer.GetString()

                utils.CompareFilePaths(path1, path2)
                return 0

            case 3:
                path, _ := fuzzConsumer.GetString()

                utils.ValidatePath(path)
                return 0

            case 4:
                path, _ := fuzzConsumer.GetString()

                utils.ValidateFileExt(path)
                return 0

            case 5:
                path, _ := fuzzConsumer.GetString()

                utils.ScanFilesInDir(path)
                return 0

            case 6:
                path, _ := fuzzConsumer.GetString()

                utils.IsDir(path)
                return 0

            case 7:
                pattern, _ := fuzzConsumer.GetString()

                utils.ValidateRegexPattern(pattern)
                return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}