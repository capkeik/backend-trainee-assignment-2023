package static

import (
	"encoding/csv"
	"os"
)

type CSVRepo struct {
	csvDir string
}

func NewCSV(dir string) CSVRepo {
	return CSVRepo{csvDir: dir}
}

func (r *CSVRepo) SaveCSV(filename string, data [][]string) error {
	filePath := r.csvDir + filename

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)

	for _, row := range data {
		err := csvWriter.Write(row)
		if err != nil {
			return err
		}
	}
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return err
	}
	return nil
}

func (r *CSVRepo) GetCSV(filename string) (os.FileInfo, string, error) {
	filePath := r.csvDir + filename

	fileInfo, err := os.Stat(filePath)

	if err != nil {
		return nil, "", err
	}

	return fileInfo, filePath, nil
}
