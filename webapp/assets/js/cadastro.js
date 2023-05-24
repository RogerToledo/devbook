$('#form-cadastro').on('submit', criarUsuario)

function criarUsuario(evento) {
    evento.preventDefault()
    console.log("Dentro da funcao")

    if ($('#senha').val() != $('#confirmaSenha').val() ) {
        Swal.fire(
            'Ooops!',
            'As senhas são diferentes!',
            'fail'
        )
        return
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
            senha: $('#senha').val()
        }
    }).done(function() {
        Swal.fire({
            title: 'Sucesso!',
            text: 'Usuário cadastrado com sucesso!',
            icon: 'success'
        }).then(function() {
            $.ajax({
                url: '/login',
                method: 'POST',
                data: {
                    email: $('#email').val(),
                    senha: $('#senha').val()
                }
            }).done(function() {
                window.location = '/home';
            }).fail(function() {
                Swal.fire({
                    title: 'Erro!',
                    text: 'Erro ao logar!',
                    icon: 'error'
                });
            });
        });    
    }).fail(function() {
        Swal.fire({
            title: 'Erro!',
            text: 'Erro ao cadastrado usuário!',
            icon: 'error'
        });
    })
}