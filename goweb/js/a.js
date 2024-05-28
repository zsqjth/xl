$(document).ready(function() {
    // 初始化商品列表
    loadProducts();

    // 搜索商品
    $('#searchButton').on('click', function(e) {
        e.preventDefault();
        var searchQuery = $('#searchBox').val();
        searchProducts(searchQuery);
    });

    // 搜索商品的函数
    function searchProducts(query) {
        $.ajax({
            url: 'http://localhost:8080/product/show?name=' + encodeURIComponent(query),
            type: 'GET',
            success: function(response) {
                // 检查 response.code 和 response.data 是否有效
                if(response.code === 0 && response.data) {
                    // 搜索返回单个 product 结构体，而不是数组
                    var product = response.data;
                    // 检查 product 是否包含必要的属性
                    if(product && product.name && product.img && product.price) {
                        // 将单个 product 放入数组中传递给 updateProductList
                        updateProductList([product]);
                    } else {
                        $('#productList').html('<p>搜索结果无效。</p>');
                    }
                } else {
                    $('#productList').html('<p>没有找到商品。</p>');
                }
            },
            error: function(xhr) {
                $('#productList').html('<p>搜索商品时发生错误。</p>');
            }
        });
    }

    // 添加商品的模态窗口逻辑
    $('#addButton').on('click', function() {
        $('#addProductModal').modal('show');
    });

     // 打开文件上传模态框的事件绑定
     $('#uploadButton').on('click', function() {
        $('#uploadModal').modal('show');
    });

    // 文件上传表单提交事件
    $('#uploadForm').on('submit', function(e) {
        e.preventDefault();
        var formData = new FormData(this);
        $.ajax({
            url: 'http://localhost:8080/upload',
            type: 'POST',
            data: formData,
            contentType: false,
            processData: false,
            xhr: function() {
                var xhr = $.ajaxSettings.xhr();
                if (xhr.upload) {
                    xhr.upload.addEventListener('progress', function(event) {
                        if (event.lengthComputable) {
                            var percentComplete = (event.loaded / event.total) * 100;
                            $('#uploadMsg').text('上传进度: ' + percentComplete.toFixed(2) + '%');
                        }
                    }, false);
                }
                return xhr;
            },
            success: function(response) {
                if(response.code === 0 && response.data && response.data.filename) {
                    $('#uploadMsg').text('文件上传成功: ' + response.data.filename);
                    $('#productImg').val(response.data.filename);
                    $('#addProductModal').modal('show');
                } else {
                    $('#uploadMsg').text('文件上传失败: ' + (response.msg || ''));
                }
            },
            error: function(xhr) {
                $('#uploadMsg').text('文件上传发生错误');
            }
        });
    });

    // 添加商品表单提交事件
    $('#addProductForm').on('submit', function(e) {
        e.preventDefault();
        var productData = {
            name: $('#productName').val(),
            img: $('#productImg').val(),
            price: $('#productPrice').val()
        };
        $.ajax({
            url: 'http://localhost:8080/product/insert',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(productData),
            success: function(response) {
                if(response.code === 0) {
                    loadProducts();
                } else {
                    alert('添加商品失败：' + response.msg);
                }
            },
            error: function(xhr) {
                alert('添加商品时发生错误。');
            }
        });
    });

    // 删除商品的逻辑
    $('#productList').on('click', '.delete-product', function() {
        var productName = $(this).data('product-name');
        if (confirm('确定要删除该商品吗？')) {
            $.ajax({
                url: 'http://localhost:8080/product/destory',
                type: 'DELETE',
                contentType: 'application/json',
                data: JSON.stringify({ name: productName }),
                success: function(response) {
                    if(response.code === 0) {
                        loadProducts(); // 刷新商品列表
                    } else {
                        alert('删除商品失败：' + response.msg);
                    }
                },
                error: function(xhr) {
                    alert('删除商品时发生错误。');
                }
            });
        }
    });

    // 加载商品列表
    function loadProducts() {
        $.ajax({
            url: 'http://localhost:8080/product/showall',
            type: 'GET',
            success: function(response) {
                if(response.code === 0 && Array.isArray(response.data)) {
                    updateProductList(response.data);
                } else {
                    $('#productList').html('<p>无法加载商品列表。</p>');
                }
            },
            error: function(xhr) {
                $('#productList').html('<p>加载商品列表时发生错误。</p>');
            }
        });
    }

    // 更新商品列表的辅助函数
    function updateProductList(products) {
        var listHtml = '';
        if (products.length === 0) {
            listHtml += '<p>没有商品信息。</p>';
        } else {
            products.forEach(function(product) {
                listHtml += '<div class="card mb-3" style="width: 18rem;">';
                listHtml += '<img src="' + product.img + '" class="card-img-top" alt="商品图片">';
                listHtml += '<div class="card-body">';
                listHtml += '<h5 class="card-title">' + product.name + '</h5>';
                listHtml += '<p class="card-text">价格: ' + product.price + '</p>';
                listHtml += '<button class="btn btn-danger delete-product" data-product-name="' + product.name + '">删除</button>';
                listHtml += '</div>';
                listHtml += '</div>';
            });
        }
        $('#productList').html(listHtml);
    }
});