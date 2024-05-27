$(document).ready(function() {
    // 初始化商品列表显示
    loadProducts();

    // 绑定搜索按钮的点击事件
    $('#searchButton').on('click', function(e) {
        e.preventDefault();
        var searchQuery = $('#searchBox').val();
        searchProducts(searchQuery);
    });

    // 绑定添加按钮的点击事件
    $('#addButton').on('click', function() {
        $('#uploadModal').modal('show');
    });

    // 处理文件上传
    $('#uploadFile').on('change', function() {
        var formData = new FormData();
        formData.append('file', $('#uploadFile')[0].files[0]);

        $.ajax({
            url: 'http://localhost:8080/upload', // 您的文件上传API端点
            type: 'POST',
            data: formData,
            contentType: false,
            processData: false,
            success: function(response) {
                if(response.code === 0 && response.data && response.data.filename) {
                    $('#uploadMsg').text('文件上传成功: ' + response.data.filename);
                    $('#productImg').val(response.data.filename); // 设置商品图片字段的值
                    $('#addProductModal').modal('show'); // 显示添加商品模态框
                } else {
                    $('#uploadMsg').text('文件上传失败: ' + (response.msg || ''));
                }
            },
            error: function(xhr) {
                $('#uploadMsg').text('文件上传发生错误');
            }
        });
    });

    // 处理添加商品表单的提交事件
    $('#addProductForm').on('submit', function(e) {
        e.preventDefault();
        var productData = {
            name: $('#productName').val(),
            img: $('#productImg').val(),
            price: $('#productPrice').val()
        };
        $.ajax({
            url: 'http://localhost:8080/product/insert', // 您的添加商品API端点
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(productData),
            success: function(response) {
                if(response.code === 0) {
                    $('#addProductModal').modal('hide');
                    loadProducts(); // 刷新商品列表
                } else {
                    alert('添加商品失败：' + (response.msg || ''));
                }
            },
            error: function(xhr) {
                alert('添加商品时发生错误。');
            }
        });
    });

    // 绑定删除商品按钮的点击事件
    $('#productList').on('click', '.delete-product', function() {
        var productName = $(this).data('product-name');
        if (confirm('确定要删除该商品吗？')) {
            $.ajax({
                url: 'http://localhost:8080/product/destory', // 您的删除商品API端点
                type: 'DELETE',
                contentType: 'application/json',
                data: JSON.stringify({ name: productName }),
                success: function(response) {
                    if(response.code === 0) {
                        loadProducts(); // 刷新商品列表
                    } else {
                        alert('删除商品失败：' + (response.msg || ''));
                    }
                },
                error: function(xhr) {
                    alert('删除商品时发生错误。');
                }
            });
        }
    });

    // 加载商品列表的函数
    function loadProducts() {
        $.ajax({
            url: 'http://localhost:8080/product/showall', // 获取商品列表的API端点
            type: 'GET',
            success: function(response) {
                if(response.code === 0 && Array.isArray(response.data)) {
                    var products = response.data;
                    var listHtml = products.map(function(product) {
                        return (
                            '<div class="card mb-3" style="width: 18rem;">' +
                            '<img src="' + product.img + '" class="card-img-top" alt="商品图片">' +
                            '<div class="card-body">' +
                            '<h5 class="card-title">' + product.name + '</h5>' +
                            '<p class="card-text">价格: ' + product.price + '</p>' +
                            '<button class="btn btn-danger delete-product" data-product-name="' + product.name + 
                            '">删除</button>' +
                            '</div>' +
                            '</div>'
                        );
                    }).join('');
                    $('#productList').html(listHtml);
                } else {
                    $('#productList').html('<p>没有商品信息。</p>');
                }
            },
            error: function(xhr) {
                $('#productList').html('<p>无法加载商品列表。</p>');
            }
        });
    }

    // 搜索商品的函数
    function searchProducts(searchQuery) {
        // 这里应实现搜索逻辑，并通过AJAX请求后端API
        // 以下为示例代码，应根据实际情况进行调整
        var product = { name: searchQuery, img: 'path/to/default/image.jpg', price: 0 };
        displayProduct(product);
    }

    // 显示单个商品信息的函数
    function displayProduct(product) {
        var listHtml = '<div class="card mb-3" style="width: 18rem;">' +
                       '<img src="' + product.img + '" class="card-img-top" alt="商品图片">' +
                       '<div class="card-body">' +
                       '<h5 class="card-title">' + product.name + '</h5>' +
                       '<p class="card-text">价格: ' + product.price + '</p>' +
                       '</div>' +
                       '</div>';
        $('#productList').html(listHtml);
    }
});