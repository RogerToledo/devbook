$('#nova-publicacao').on('submit', criarPublicacao);
$('.curtir-publicacao').on('click', curtirPublicacao);

$('#atualizar-publicacao').on('click', atualizarPublicacao)
$('.deletar-publicacao').on('click', deletarPublicacao)

function criarPublicacao(evento) {
    evento.preventDefault();
    
    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function() {
        Swal.fire({
            title: 'Erro!',
            text: 'Erro ao criar publicação!',
            icon: 'error'
        });
    });
}

function atualizarPublicacao(evento) {
    $(this).prop('disable', false);

    const publicacaoID = $(this).data('publicacao-id');

    $.ajax({
        url: `/publicacoes/editar/${publicacaoID}`,
        method: "PUT",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        }
    }).done(function() {
        Swal.fire(
            'Sucesso!',
            'Publicação atualizada com sucesso!',
            'success'
        ).then(function() {
            window.location = '/home'
        });
    }).fail(function(e) {
        console.error(e)
        Swal.fire({
            title: 'Erro!',
            text: 'Erro ao salvar publicação!',
            icon: 'error'
        });
    }).always(function() {
        $('#atualizar-publicacao').prop('disabled', true);
    });
}

function deletarPublicacao(evento) {
    Swal.fire({
        title: 'Tem cereza??',
        text: 'Essa ação não pode ser desfeita.',
        showCancelButton: true,
        cancelButtonText: 'Cancela',
        icon: 'warning'
    }).then(function(confirmacao) {
        if (!confirmacao.value) return;

        evento.preventDefault();
        const elementoClicado = $(evento.target);
        const publicacao = elementoClicado.closest('div')
        const publicacaoID = publicacao.data('publicacao-id');
        
        elementoClicado.prop('disable', true); 

        $.ajax({
            url: `/publicacoes/${publicacaoID}`,
            method: "DELETE"
        }).done(function() {
            publicacao.fadeOut("slow", function() {
                $(this).remove();
            });
        }).fail(function() {
            Swal.fire({
                title: 'Erro!',
                text: 'Erro ao deletar publicação!',
                icon: 'error'
            });
        });
    });
}

function curtirPublicacao(evento) {
    evento.preventDefault();
    const elementoClicado = $(evento.target);
    const publicacaoID = elementoClicado.closest('div').data('publicacao-id');
    
    elementoClicado.prop('disable', true);

    $.ajax({
        url: `/publicacoes/curtir/${publicacaoID}`,
        method: "POST",
    }).done(function() {
        window.location = "/home";
        elementoClicado.addClass('text-danger')
    }).fail(function() {
        Swal.fire({
            title: 'Erro!',
            text: 'Erro ao curtir publicação!',
            icon: 'error'
        });
    }).always(function() {
        elementoClicado.prop('disable', false);
    });
}
