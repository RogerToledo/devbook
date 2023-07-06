$('#parar-seguir').on('click', pararSeguir)
$('#seguir').on('click', seguir)
$('#editar-usuario').on('submit', editarUsuario)
$('#atualizar-senha').on('submit', atualizarSenha)
$('#deletar-usuario').on('click', deletarUsuario)

function pararSeguir() {
    const usuarioID = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/parar-seguir/${usuarioID}`,
        method: "POST"
    }).done(function() {
        window.location = `/usuarios/${usuarioID}`;
    }).fail(function() {
        Swal.fire({
            title: 'Ops...',
            text: 'Erro ao parar de seguir usuário!',
            icon: 'error'
        });
        $('#parar-seguir').prop('disabled', false);
    })
}

function seguir() {
    const usuarioID = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/seguir/${usuarioID}`,
        method: "POST"
    }).done(function() {
        window.location = `/usuarios/${usuarioID}`;
    }).fail(function() {
        Swal.fire({
            title: 'Ops...',
            text: 'Erro ao parar de seguir usuário!',
            icon: 'error'
        });
        $('#seguir').prop('disabled', false);
    })
}

function editarUsuario(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/editar-usuario",
        method: "PUT",
        data: {
            nome: $('#nome').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
        }
    }).done(function() {
        Swal.fire({
            title: 'Sucesso!',
            text: 'Usuário editado com sucesso!',
            icon: 'success'
        }).then(function() {
            window.location = '/perfil';
        })
    }).fail(function() {
        Swal.fire({
            title: 'Ops...',
            text: 'Erro ao editar usuário!',
            icon: 'error'
        })
    })
}

function atualizarSenha(evento) {
    evento.preventDefault();

    if ($('#nova-senha').val() != $('#confirmar-senha').val()) {
        Swal.fire({
            title: 'Ops...',
            text: 'As senhas não conferem!',
            icon: 'error'
        })
        return;
    }

    $.ajax({
        url: "/atualizar-senha",
        method: "POST",
        data: {
            senhaAtual: $('#senha-atual').val(),
            novaSenha: $('#nova-senha').val(),
          }
    }).done(function() {
        Swal.fire({
            title: 'Sucesso!',
            text: 'Senha atualizada com sucesso!',
            icon: 'success'
        }).then(function() {
            window.location = '/perfil';
        })
    }).fail(function() {
        Swal.fire({
            title: 'Ops...',
            text: 'Erro ao atualizar senha!',
            icon: 'error'
        })
    })
}

function deletarUsuario() {
    Swal.fire({
        title: 'Atenção!',
        text: 'Tem certeza que deseja excluir sua conta? Essa ação é irreversível!',
        showCancelButton: true,
        icon: 'warning',
    }).then(function(confirmacao) {
        if (confirmacao.value) {
            $.ajax({
                url: "/deletar-usuario",
                method: "DELETE"
            }).done(function() {
                Swal.fire({
                    title: 'Sucesso!',
                    text: 'Usuário deletado com sucesso!',
                    icon: 'success'
                }).then(function() {
                    window.location = '/logout';
                })
            }).fail(function() {
                Swal.fire({
                    title: 'Ops...',
                    text: 'Erro ao deletar usuário!',
                    icon: 'error'
                })
            })
        }
    })
}