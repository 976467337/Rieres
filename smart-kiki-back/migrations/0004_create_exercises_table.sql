-- +goose Up
CREATE TABLE exercises (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    muscle_group VARCHAR(50) NOT NULL,
    equipment VARCHAR(100) NOT NULL DEFAULT '',
    instructions TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_exercises_muscle_group ON exercises (muscle_group);

INSERT INTO exercises (name, muscle_group, equipment, instructions) VALUES
('Supino reto', 'Peito', 'Barra', 'Deite no banco, desça a barra até o peito e empurre de volta.'),
('Supino inclinado', 'Peito', 'Barra', 'Banco inclinado a 30-45 graus, mesmo movimento do supino reto.'),
('Supino declinado', 'Peito', 'Barra', 'Banco declinado, foco na porção inferior do peitoral.'),
('Crucifixo', 'Peito', 'Halteres', 'Braços semiflexionados, abra e feche em arco sobre o peito.'),
('Crossover', 'Peito', 'Cabo', 'Puxe os cabos das polias altas em direção ao centro do corpo.'),
('Flexão de braço', 'Peito', 'Peso corporal', 'Apoie mãos e pés no chão, flexione os cotovelos até quase tocar o chão.'),
('Puxada frontal', 'Costas', 'Cabo', 'Puxe a barra até a altura do queixo, cotovelos apontando pra baixo.'),
('Remada baixa', 'Costas', 'Cabo', 'Sentado, puxe o cabo em direção ao abdômen mantendo a coluna reta.'),
('Remada curvada', 'Costas', 'Barra', 'Tronco inclinado à frente, puxe a barra em direção ao abdômen.'),
('Barra fixa', 'Costas', 'Peso corporal', 'Suspenso na barra, puxe o corpo até o queixo passar da barra.'),
('Pulldown', 'Costas', 'Cabo', 'Similar à puxada frontal, com pegada mais aberta.'),
('Levantamento terra', 'Costas', 'Barra', 'Levante a barra do chão mantendo a coluna neutra até a extensão do quadril.'),
('Agachamento livre', 'Pernas', 'Barra', 'Barra nas costas, agache mantendo o tronco ereto até coxas paralelas ao chão.'),
('Leg press', 'Pernas', 'Máquina', 'Empurre a plataforma com os pés, controlando a descida.'),
('Cadeira extensora', 'Pernas', 'Máquina', 'Sentado, estenda os joelhos contra a resistência.'),
('Mesa flexora', 'Pernas', 'Máquina', 'Deitado de bruços, flexione os joelhos trazendo os calcanhares aos glúteos.'),
('Afundo', 'Pernas', 'Halteres', 'Passo à frente, desça até o joelho de trás quase tocar o chão.'),
('Cadeira adutora', 'Pernas', 'Máquina', 'Sentado, feche as pernas contra a resistência.'),
('Cadeira abdutora', 'Pernas', 'Máquina', 'Sentado, abra as pernas contra a resistência.'),
('Panturrilha em pé', 'Pernas', 'Máquina', 'Eleve os calcanhares o máximo possível, contraindo a panturrilha.'),
('Desenvolvimento militar', 'Ombro', 'Barra', 'Empurre a barra acima da cabeça a partir dos ombros.'),
('Elevação lateral', 'Ombro', 'Halteres', 'Eleve os braços lateralmente até a altura dos ombros.'),
('Elevação frontal', 'Ombro', 'Halteres', 'Eleve os braços à frente até a altura dos ombros.'),
('Remada alta', 'Ombro', 'Barra', 'Puxe a barra verticalmente até a altura do peito, cotovelos altos.'),
('Encolhimento', 'Ombro', 'Halteres', 'Eleve os ombros em direção às orelhas e controle a descida.'),
('Rosca direta', 'Bíceps', 'Barra', 'Flexione os cotovelos elevando a barra até os ombros.'),
('Rosca alternada', 'Bíceps', 'Halteres', 'Flexione um braço de cada vez, alternando.'),
('Rosca martelo', 'Bíceps', 'Halteres', 'Pegada neutra, flexione os cotovelos sem girar o pulso.'),
('Rosca scott', 'Bíceps', 'Barra', 'Apoiado no banco scott, flexione os cotovelos isolando o bíceps.'),
('Tríceps pulley', 'Tríceps', 'Cabo', 'Empurre a barra/corda para baixo estendendo os cotovelos.'),
('Tríceps testa', 'Tríceps', 'Barra', 'Deitado, desça a barra em direção à testa flexionando os cotovelos.'),
('Tríceps corda', 'Tríceps', 'Cabo', 'Puxe a corda para baixo abrindo as mãos ao final do movimento.'),
('Mergulho no banco', 'Tríceps', 'Peso corporal', 'Mãos no banco atrás do corpo, flexione e estenda os cotovelos.'),
('Abdominal supra', 'Core', 'Peso corporal', 'Deitado, eleve o tronco em direção aos joelhos.'),
('Prancha', 'Core', 'Peso corporal', 'Apoie antebraços e pés, mantenha o corpo reto e contraia o abdômen.'),
('Elevação de pernas', 'Core', 'Peso corporal', 'Deitado, eleve as pernas estendidas até 90 graus.'),
('Abdominal oblíquo', 'Core', 'Peso corporal', 'Gire o tronco em direção aos joelhos alternando os lados.'),
('Prancha lateral', 'Core', 'Peso corporal', 'Apoiado de lado em um antebraço, mantenha o quadril elevado.'),
('Ab wheel', 'Core', 'Roda abdominal', 'Ajoelhado, role a roda à frente e retorne controlando o core.'),
('Esteira', 'Cardio', 'Esteira', 'Caminhada ou corrida em ritmo constante ou intervalado.'),
('Bicicleta ergométrica', 'Cardio', 'Bicicleta', 'Pedalada em ritmo constante ou intervalado.'),
('Corda naval', 'Cardio', 'Corda naval', 'Ondulações alternadas ou simultâneas com as cordas.'),
('Burpee', 'Cardio', 'Peso corporal', 'Agache, jogue as pernas para trás, flexão, salte.'),
('Polichinelo', 'Cardio', 'Peso corporal', 'Salte abrindo pernas e braços simultaneamente.'),
('Escada de agilidade', 'Cardio', 'Escada de agilidade', 'Passadas rápidas e coordenadas entre os degraus da escada.');

-- +goose Down
DROP TABLE IF EXISTS exercises;
