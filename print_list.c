#include "common.h"

void print_list(node *head)
{
        /* Check if list is empty */
        if (head == NULL)
                printf("Failed: Sorry, List is Empty\n");
        else
        {
                /* traverse and print all the node datas */
                while (head)
                {
#if 1
                        if (head->data < 10)
                                printf("000");
                        else if (head->data > 9 && head->data < 100)
                                printf("00");
                        else if (head->data > 99 && head->data < 1000)
                                printf("0");
#endif
                        printf("%d ", head->data);
                        head = head->next;
                }
                putchar('\n');
        }
}

void print_back(node *tail)
{
        if (tail == NULL)
                printf("List is empty\n");
        else
        {
                while (tail)
                {
                        printf("%d ", tail->data);
                        tail = tail->prev;
                }
                putchar('\n');
        }
}
