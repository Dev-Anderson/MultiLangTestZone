import pandas as pd

# carregando as planilhas
planilha01 = pd.read_csv('05-08-2024.csv') 
planilha02 = pd.read_csv('block-05-08-2024.csv')

# filtrando apenas as coisas do Gestao de Frotas
filtro_frotas = planilha01[planilha01['historical_board'] == 'Gestão de Frotas']
card_ids_planilha2 = planilha02['card_id']

# pegando o numero do card e vendo se tem dentro da planilha 01 
filtro_frotas_sem_planilha2 = filtro_frotas[~filtro_frotas['card_id'].isin(card_ids_planilha2)]

qtde_registro = filtro_frotas_sem_planilha2.shape[0]
soma_size = filtro_frotas_sem_planilha2['size'].sum()
qtde_block = planilha02['card_id'].shape[0]
soma_size_block = planilha02['size'].sum()

# salvando dentro do csv 
resultado = pd.DataFrame({'Qtde tarefas': [qtde_registro], 'Total pontos': [soma_size], 'Qtde tarefas block': [qtde_block], 'Total pontos block': [soma_size_block]})
resultado.to_csv('resultado.csv', index=False)
