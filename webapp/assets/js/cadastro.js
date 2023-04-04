$('#form-cadastro').on('submit', criarUsuario)

function criarUsuario(evento) {
    evento.preventDefault()
    console.log("Dentro da funcao")

    if ($('#senha').val() != $('#confirmaSenha').val() ) {
        alert("As senhas são diferentes!")
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
    })
}