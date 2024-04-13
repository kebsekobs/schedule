import React from 'react';
import { useForm } from 'react-hook-form';
import {useAddGroupsMutation} from "../api/AddGroupMutation";

const AddGroupModal = ({ isOpen, toggleModal }) => {
    const form = useForm();
    const { reset, register, handleSubmit, formState: { errors } } = form;

    const addGroupMutation = useAddGroupsMutation();
    const onSubmit = data => {
        console.log(data)
        addGroupMutation.mutateAsync(data)
        toggleModal();
        reset()
    };

    if (!isOpen) {
        return null;
    }

    return (
        <div style={styles.backdrop}>
            <div style={styles.modal}>
                <h2>Введите данные</h2>
                <form onSubmit={handleSubmit(onSubmit)}>
                    <input
                        type={"number"}
                        {...register("groupId", { required: true })}
                        placeholder="Введите номер группы"
                        style={styles.input}
                    />
                    {errors.inputId && <span style={styles.error}>Это поле обязательно</span>}
                    <input
                        {...register("name", { required: true })}
                        placeholder="Введите название группы"
                        style={styles.input}
                    />
                    {errors.inputName && <span style={styles.error}>Это поле обязательно</span>}
                    <input
                        type={"number"}
                        {...register("capacity", { required: true })}
                        placeholder="Введите количество студентов"
                        style={styles.input}
                    />
                    {errors.inputCapacity && <span style={styles.error}>Это поле обязательно</span>}
                    <label htmlFor="consent" style={styles.label}>
                        <input
                            type="checkbox"
                            id="magistracy"
                            style={styles.checkbox}
                            {...register("magistracy")}
                        />
                       Магистратура
                    </label>
                    <button type="submit" style={styles.button}>Отправить</button>
                </form>
                <button onClick={toggleModal} style={styles.buttonClose}>Закрыть</button>
            </div>
        </div>
    );
};

const styles = {
    backdrop: {
        position: 'fixed',
        top: 0,
        left: 0,
        right: 0,
        bottom: 0,
        backgroundColor: 'rgba(0,0,0,0.5)',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        zIndex: 1000,
    },
    modal: {
        backgroundColor: '#fff',
        padding: '20px',
        borderRadius: '8px',
        display: 'flex',
        flexDirection: 'column',
        gap: '10px',
    },
    input: {
        width: '300px',
        padding: '10px',
        margin: '10px 0',
        borderRadius: '4px',
        border: '1px solid #ccc',
    },
    checkbox: {
        width: '20px',
        height: '20px',
        margin: '10px 5px',
        cursor: 'pointer'
    },
    button: {
        padding: '10px 20px',
        border: 'none',
        backgroundColor: 'blue',
        color: 'white',
        borderRadius: '5px',
        cursor: 'pointer',
        marginTop: '10px'
    },
    buttonClose: {
        padding: '10px 20px',
        border: 'none',
        backgroundColor: 'gray',
        color: 'white',
        borderRadius: '5px',
        cursor: 'pointer',
        marginTop: '10px'
    },
    error: {
        color: 'red',
        fontSize: '14px',
        marginTop: '5px'
    },
    label: {
        display: 'flex',
        alignItems: 'center',
        cursor: 'pointer'
    }
};

export default AddGroupModal;
