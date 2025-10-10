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

## Классы эквивалентности
### Конечность классов эквивалентности по НФ
Заметим, что если слева стоит $b$,  то вся строка приводитя к виду $b^n$. Следовательно, нужно рассмотреть два набора:
1. набор, где невозможно привести $b$ влево;
2. последовательности $b^n$.

Рассмотрим первый набор:

$$a, aa, ab, aab, aba, abb, aaba, aabb, abbb, aabbb$$

Следует обратить внимание на $aaa$:

$$aaa \leftrightarrow abab \leftrightarrow ab^5 \leftrightarrow abba \leftrightarrow baaab \leftrightarrow bababb \leftrightarrow b^{10} \leftrightarrow b^6ba \leftrightarrow b^3baa \leftrightarrow baaa$$

Таким образом, мы получаем 

$$aaa \leftrightarrow baaa$$

$$aaa \leftrightarrow b^{10}$$

Аналогично, 

$$aaaa \leftrightarrow ababa \leftrightarrow ab^5a \leftrightarrow abbaa \leftrightarrow baaaba \leftrightarrow bababba \leftrightarrow b^{13} \leftrightarrow b^3aaa \leftrightarrow aaa$$

То есть, 

$$aaa \leftrightarrow aaaa$$

Рассмотрим второй набор:

1. $b$
2. $bb$
3. $bbb$
4. $$bbbb \leftrightarrow ba$$
5. $$bbbbb \leftrightarrow bab$$
6. $$bbbbbb \leftrightarrow babb$$
7. $$b^7 \leftrightarrow b^4a \leftrightarrow baa $$
8. $$b^8 \leftrightarrow baab $$
9. $$b^9 \leftrightarrow b^4b^4b \leftrightarrow babab \leftrightarrow baaa \leftrightarrow aaa $$
10. $$b^{10} \leftrightarrow baaa \leftrightarrow aaa $$

При $$n \geq 9$$  $$b^n \leftrightarrow  aaa$$.

Мы получаем классы эквивалентности: $$\epsilon, a, aa, ab, aab, aba, abb, aaba, aabb, abbb, aabbb, b, bb, bbb, ba, bab, babb, baa, baab, aaa$$.


### Построение ДКА
Построим ДКА по данным классам эквивалентности.
![ДКА](img/automa.svg)

### Построение минимальной системы переписывания
Воспользовавашись опытом прошлых цивилизаций, а точнее работой ["Построение трансформационного моноида по автомату (А.Иванов)](https://github.com/UsefulTornado/Formal-Languages), мы получили минимальную систему переписывания:

$$
bba \rightarrow bab \\
$$
$$
aaaa \rightarrow aaa \\
$$
$$
aaab \rightarrow aaa \\
$$
$$
abaa \rightarrow aaa \\
$$
$$
abab \rightarrow aaa \\
$$
$$
baaa \rightarrow aaa \\
$$
$$
baba \rightarrow baab \\
$$
$$
bbbb \rightarrow ba \\
$$
$$
baaba \rightarrow aaa \\
$$
$$
baabb \rightarrow aaa \\
$$
$$
babbb \rightarrow baa \\
$$


## Локальная конфлюэнтность
Данная система переписывания $\tau$ незавершима, следовательно, не существует фундированной монотонной алгебры, совместимой с этой системой.
### Переупорядочивание правил
Переориентируем правила, чтобы задать фундированный порядок.

$$
aabbbaa \rightarrow abbaba \\
$$
$$
abab \rightarrow aaa \\
$$
$$
baaab \rightarrow abba \\ 
$$
$$
bbbb \rightarrow ba \\
$$

Теперь SRS завершима, так как каждое правило уменьшает длину слова.

Фундированный порядок: shortlex (армейский), где $$a < b$$. 
### Проверка конфлюэнтности
Рассмотрим слово $bbbbb$:

$$bbbbb \rightarrow bab$$

$$bbbbb \rightarrow bba$$

Как мы видим, SRS не конфлюэнтна.

### Пополняемость по Кнуту-Бендикса
Доведем сиситему до конфлюэнтности. 

1. Из предыдущего пункта получаем:

$$bba \rightarrow bab$$

2. Проверим слово $bbbba$:

$$bbbba \rightarrow baa$$

$$bbbba \rightarrow bbbab \rightarrow bbabb \rightarrow babbb$$

Получаем: $$babbb \rightarrow baa$$

3. Проверим слово $babbbb$:

$$babbbb \rightarrow baab$$

$$babbbb \rightarrow baba$$

Получаем: $$baba \rightarrow baab$$

4. Проверим слово $babbba$:

$$babbba \rightarrow baaa$$

$$babbba \rightarrow babbab \rightarrow bababb \rightarrow baaab \rightarrow abba \rightarrow abab \rightarrow aaa$$

Получаем: $$baaa \rightarrow aaa$$

5. Проверим слово $baaab$:

$$baaab \rightarrow aaab$$

$$baaab \rightarrow abba \rightarrow abab \rightarrow aaa$$

Получаем: $$aaab \rightarrow aaa$$

6. Проверим слово $babab$: 

$$babab \rightarrow baabb$$

$$babab \rightarrow baaaa \rightarrow aaa$$

Получаем: $$baabb \rightarrow aaa$$

7. Проверим слово $abbbbba$: 

$$abbbbba \rightarrow ababa \rightarrow aaaa$$

$$abbbbba \rightarrow abbbbab \rightarrow abbbabb \rightarrow abbabbb \rightarrow ababbbb \rightarrow aaabbbb \rightarrow aaa$$

Получаем: $$aaaa \rightarrow aaa$$

7. Ещё раз рассмотрим слово $abbbbba$: 

$$abbbbba \rightarrow ababa \rightarrow baaaba$$

$$abbbbba \rightarrow abbbbab \rightarrow abbbabb \rightarrow abbabbb \rightarrow ababbbb \rightarrow aaabbb \rightarrow aaa$$

Получаем: $$baaaba \rightarrow aaa$$

8. Проверим слово $abbbba$: 

$$abbbba \rightarrow abaa$$

$$abbbba \rightarrow abbbab \rightarrow abbabb \rightarrow ababbb \rightarrow aaabb \rightarrow aaa$$

Получаем: $$abaa \rightarrow aaa$$


Пополненая система:

$$
aabbbaa \rightarrow abbaba \\
$$
$$
abab \rightarrow aaa \\
$$
$$
baaab \rightarrow abba \\ 
$$
$$
bbbb \rightarrow ba \\
$$
$$
bba \rightarrow bab \\
$$
$$
aaaa \rightarrow aaa \\
$$
$$
aaab \rightarrow aaa \\
$$
$$
abaa \rightarrow aaa \\
$$
$$
baaa \rightarrow aaa \\
$$
$$
baba \rightarrow baab \\
$$
$$
baaba \rightarrow aaa \\
$$
$$
baabb \rightarrow aaa \\
$$
$$
babbb \rightarrow baa \\
$$


Заметим, что при пополнение по Кнуту-Бендиксу мы получаем те же правила, что были в минимальной системе. Это объясняется тем, что из конечного числа классов, следует конечное множество переходов между ними.


## Инварианты 
Исходная SRS $\tau$ (с переориентированными правилами):

$$
aabbbaa \rightarrow abbaba \\
$$
$$
abab \rightarrow aaa \\
$$
$$
baaab \rightarrow abba \\ 
$$
$$
bbbb \rightarrow ba \\
$$

Минимальная SRS $\tau'$:

$$
bba \rightarrow bab \\
$$
$$
aaaa \rightarrow aaa \\
$$
$$
aaab \rightarrow aaa \\
$$
$$
abaa \rightarrow aaa \\
$$
$$
abab \rightarrow aaa \\
$$
$$
baaa \rightarrow aaa \\
$$
$$
baba \rightarrow baab \\
$$
$$
bbbb \rightarrow ba \\
$$
$$
baaba \rightarrow aaa \\
$$
$$
baabb \rightarrow aaa \\
$$
$$
babbb \rightarrow baa \\
$$

Инварианты:
1. Всегда есть буква $a$.
2. Если есть $bbb$, то оно сокращается.