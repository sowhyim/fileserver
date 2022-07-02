function Download(filename, uri, location, content) {
    axios.post(uri, {}, {
        responseType: 'blob'
    }).then(function (res) {
        let blob = res.data;
        // FileReader主要用于将文件内容读入内存
        let reader = new FileReader();
        reader.readAsDataURL(blob);
        // onload当读取操作成功完成时调用
        reader.onload = function (e) {
            let a = document.createElement('a');
            // 获取文件名fileName
            let fileName = res.headers["content-disposition"].split("=");
            fileName = fileName[fileName.length - 1];
            fileName = fileName.replace(/"/g, "");
            a.download = fileName;
            a.href = e.target.result;
            document.body.appendChild(a);
            a.click();
            document.body.removeChild(a);
        }
    });
}