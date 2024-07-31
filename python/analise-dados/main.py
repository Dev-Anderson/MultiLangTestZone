import pandas as pd 
df = pd.read_csv('30-07-2024.csv') #lendo o arquivo csv
frotas = df[df['historical_board'] == 'Gestão de Frotas'].shape[0]
filtro = (df['historical_board'] == 'Gestão de Frotas')
soma_frotas = df.loc[filtro, 'size'].sum()

# qtde backlog - backend 
backlog_back = (df['historical_board'] == 'Gestão de Frotas') & (df['historical_column'] == 'Backlog') & (df['historical_column_parent'] == 'Back - end')
count_backlog_back = df[backlog_back].shape[0]

# pontos backlog - backend 
soma_backlog_back = (df['historical_board'] == 'Gestão de Frotas') & (df['historical_column'] == 'Backlog') & (df['historical_column_parent'] == 'Back - end')
soma_backlog_back_sum = df.loc[soma_backlog_back, 'size'].sum()

# qtde backlog - frontend 
backlog_front = (df['historical_board'] == 'Gestão de Frotas') & (df['historical_column'] == 'Backlog') & (df['historical_column_parent'] == 'Front - end')
count_backlog_front = df[backlog_front].shape[0]

# pontos backlog - frontend 
soma_backlog_front = (df['historical_board'] == 'Gestão de Frotas') & (df['historical_column'] == 'Backlog') & (df['historical_column_parent'] == 'Front - end')
soma_backlog_front_sum = df.loc[soma_backlog_front, 'size'].sum()

# qtde em andamento - backend 
andamento_back = (df['historical_board'] == 'Gestão de Frotas') & (df['historical_column'] == 'Em Andamento') & (df['historical_column_parent'] == 'Back - end')
count_andamento_back = df[andamento_back].shape[0]

# pontos em andamento - backend 
soma_andamento_back = (df['historical_board'] == 'Gestão de Frotas') & (df['historical_column'] == 'Em Andamento') & (df['historical_column_parent'] == 'Back - end')
sum_andamento_back = df.loc[soma_andamento_back, 'size'].sum()

# qtde em andamento - front 
andamento_front = (df['historical_board'] == 'Gestão de Frotas') & (df['historical_column'] == 'Em Andamento') & (df['historical_column_parent'] == 'Front - end')
count_andamento_front = df[andamento_front].shape[0]

# pontos em andamento - front 
soma_andamento_front = (df['historical_board'] == 'Gestão de Frotas') & (df['historical_column'] == 'Em Andamento') & (df['historical_column_parent'] == 'Front - end')
sum_andamento_front = df.loc[soma_andamento_front, 'size'].sum()


resultado_df = pd.DataFrame({'Qtde total': [frotas],'Total pontos': [soma_frotas], 'Qtde backlog back': [count_backlog_back], 'Soma backlog back': [soma_backlog_back_sum], 'Qtde backlog front': [count_backlog_front], 'Soma backlog front': [soma_backlog_front_sum], 'Qtde Em andamento back': [count_andamento_back], 'Soma em andamento back': [sum_andamento_back], 'Qtde Em andamento front': [count_andamento_front], 'Soma Em andamento front': [soma_andamento_front]})
resultado_df.to_csv('resultado.csv', index=False)