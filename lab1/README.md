# Лабораторная работа №1
Дана SRS $\tau$:

$$
abbaba \rightarrow aabbbaa \\
$$
$$
aaa \rightarrow abab \\
$$
$$
abba \rightarrow baaab \\ 
$$
$$
bbbb \rightarrow ba \\
$$

## Завершимость 
Данная SRS незавершима. Пример: 

$$
w_1 abbaa w_2
$$
$$
w_1 baaaba w_2
$$
$$
w_1 bababba w_2
$$
$$
w_1 babbaaab w_2
$$
$$
w'_1 abbaa w'_2
$$

Где $w'_1 > w_1$ и $w'_2 > w_2$. Мы получили бесконечный цикл.
