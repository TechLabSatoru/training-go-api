package post

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"io/ioutil"
)

func Sample() int {
	return 18
}

func HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	// レスポンスとして返すデータを作成
	response := "Hello, World!"

	// レスポンスを書き込む
	fmt.Fprint(w, response)
}

/* curl -i -f -H "content-type:application/x-binary" -X POST "localhost:18080/v1/upload" --data-binary "@input/sample.tar.gz" */
func SamplePostRequest(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintln(w, "Failed to upload")
	}

    outputPath := "./output/sample.tar.gz"

	err = ioutil.WriteFile(outputPath, buf, 0644)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "File uploaded successfully")
}

func HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	// リクエストからファイルを取得
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintln(w, "File uploaded successfully01")
		http.Error(w, "Failed to retrieve file from request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 出力先ディレクトリのパス
	outputDir := "./output"

	// ファイルを保存する
	err = saveFile(file, filepath.Join(outputDir, "sample.tar.gz"))
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	// 成功レスポンスを返す
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "File uploaded successfully")
}

func saveFile(file io.Reader, path string) error {
	// 出力先ディレクトリを作成
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	// ファイルを作成
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// ファイルに書き込む
	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	return nil
}
