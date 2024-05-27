$(document).ready(function() {
    // 注册表单提交事件
    $('#registerForm').submit(function(e) {
        e.preventDefault();
        var username = $('#username').val();
        var password = $('#password').val();
        // 发送AJAX请求到后端API进行注册
        $.ajax({
            url: 'http://localhost:8080/register', // 确保这是注册的API端点
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({username: username, password: password}),
            success: function(response) {
                if(response.code === 0) {
                    alert('注册成功');
                    $('#registerForm')[0].reset(); // 清空注册表单
                } else {
                    alert('注册失败: ' + response.msg);
                }
            },
            error: function(xhr) {
                alert('注册请求失败');
            }
        });
    });

    // 登录表单提交事件
    $('#loginForm').submit(function(e) {
        e.preventDefault();
        var username = $('#loginUsername').val();
        var password = $('#loginPassword').val();
        var rememberMe = $('#rememberMe').is(':checked');
        // 发送AJAX请求到后端API进行登录
        $.ajax({
            url: 'http://localhost:8080/login', // 确保这是登录的API端点
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({username: username, password: password}),
            success: function(response) {
                if(response.code === 0) {
                    alert('登录成功');
                    if(rememberMe) {
                        localStorage.setItem('username', username);
                        localStorage.setItem('password', password);
                    }
                    window.location.href = 'products.html'; // 登录成功后跳转到商品页面
                } else {
                    alert('登录失败: ' + response.msg);
                }
            },
            error: function(xhr) {
                alert('登录请求失败');
            }
        });
    });

    // 自动登录逻辑
    autoLogin();
});

function autoLogin() {
    var username = localStorage.getItem('username');
    var password = localStorage.getItem('password');
    if(username && password) {
        // 发送AJAX请求尝试自动登录
        $.ajax({
            url: 'http://localhost:8080/login', // 确保这是登录的API端点
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({username: username, password: password}),
            success: function(response) {
                if(response.code === 0) {
                    alert('自动登录成功');
                    window.location.href = 'products.html'; // 自动登录成功后跳转到商品页面
                } else {
                    localStorage.removeItem('username');
                    localStorage.removeItem('password');
                    alert('自动登录失败，请重新登录。');
                }
            },
            error: function(xhr) {
                localStorage.removeItem('username');
                localStorage.removeItem('password');
                alert('自动登录失败，请重新登录。');
            }
        });
    }
}