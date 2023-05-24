$('#login').on('submit', fazerLogin);

function fazerLogin(evento) {
    evento.preventDefault();

    $.ajax({
        url: '/login',
        method: 'POST',
        data: {
            email: $('#email').val(),
            senha: $('#senha').val(),
        }
    }).done(function() {
        window.location = '/home';
    }).fail(function(e) {
        console.error(e)
        Swal.fire({
            title: 'Ooops...',
            text: 'Usuário ou senha inválidos',
            icon: 'warning'
        });
    });
}