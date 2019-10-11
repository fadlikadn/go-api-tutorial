const base_url = 'http://localhost:8080';

// $('#user-form').on('submit', function(e) {
//
// });

$(document).on('click', '#btn_login_action', function(e) {
    e.preventDefault();

    var $self = $(this);
    var payload = JSON.stringify({
        email: $('#inputEmail').val(),
        password: $('#inputPassword').val()
    });

    $.ajax({
        url: base_url + '/api/login',
        method: 'POST',
        data: payload,
        contentType: 'application/json'
    }).then(function(res) {
        console.log(res);
        // alert('berhasil \n' + res);
        window.location.replace(base_url+'/');
    }).catch(function(err) {
        console.log(err);
        alert('gagal');
    });
});