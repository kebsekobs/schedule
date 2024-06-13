import React from "react";
import { useForm, Controller } from "react-hook-form";
import styles from "../../shared/style/modal.module.css";
import Button from "../../../components/button";
import { useEditDisciplinesMutation } from "../api/editDisciplinesMutation";
import { useGroupsQuery } from "../../groups/api/getGroupsQuery";
import { useTeachersQuery } from "../../teachers/api/getTeachersQuery";
import { Multiselect } from "multiselect-react-dropdown";
import {CloseSvg} from "../../../components/close-svg";

const EditDisciplinesModal = ({ isOpen, toggleModal, original }) => {
  const id = original.id;
  const {
    register,
    control,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm({
    defaultValues: {
      disciplinesId: original.disciplinesId,
      name: original.name,
      teachers: original.teachers,
      hours: original.hours,
      relatedGroupsId: original.relatedGroupsId,
    },
  });

  const editDisciplinesMutation = useEditDisciplinesMutation();
  const getGroup = useGroupsQuery().data;
  const teachersQuery = useTeachersQuery();

  const options = getGroup?.map((el) => el.subRows.map((group) => `${group.groupId} ${group.name !== undefined ? group.name : '' }`));

  console.log(options);

  const onSubmit = (data) => {
    data.id = id;
    editDisciplinesMutation.mutateAsync(data);
    toggleModal();
    reset();
  };

  if (!isOpen) {
    return null;
  }


  const multiSelectStyles = {
    multiselectContainer: {
      backgroundColor: "white",
    },
    inputField: {
      margin: "5px",
    },
    chips: {
      background: "#8090bc",
    },
    optionContainer: {
      maxHeight: "160px",
      overflowY: "scroll",
    },
  };

  return (
      <div className={styles["backdrop"]}>
        <div className={styles["modal"]}>
          <div className={styles["modal-header"]}>
            <h2>Измените данные</h2>
            <Button onClick={toggleModal} styleFeature="close">
              <CloseSvg />
            </Button>
          </div>
          <form onSubmit={handleSubmit(onSubmit)} className={styles["groups-form"]}>
            <label>Измените имя дисциплины</label>
            <input
                {...register("name", { required: "Это поле обязательно" })}
                placeholder="Ин.яз"
                className={styles["input"]}
            />
            {errors.name && (
                <span className={styles.error}>{errors.name.message}</span>
            )}
            <label>Измените кол-во часов в неделю</label>
            <input
                type="number"
                {...register("hours", {
                  required: "Это поле обязательно",
                  min: { value: 1, message: "Значение должно быть больше 0" }
                })}
                placeholder="20"
                className={styles["input"]}
            />
            {errors.hours && (
                <span className={styles.error}>{errors.hours.message}</span>
            )}
            <label>Измените преподователя</label>
            <select
                {...register("teachers", { required: "Это поле обязательно" })}
                className={styles["input"]}
            >
              {teachersQuery.data.map((el, index) => (
                  <option
                      key={index}
                      value={`${el.id}`}
                  >{`Преподователь: ${el.name}`}</option>
              ))}
            </select>
            {errors.teachers && (
                <span className={styles.error}>{errors.teachers.message}</span>
            )}
            <label>Измените группу(ы)</label>
            <Controller
                control={control}
                name="relatedGroupsId"
                render={({ field: { value, onChange } }) => (
                    <Multiselect
                        placeholder="314 GTH-JDO-NSK"
                        options={options}
                        isObject={false}
                        showCheckbox={true}
                        hidePlaceholder={true}
                        closeOnSelect={false}
                        onSelect={onChange}
                        onRemove={onChange}
                        selectedValues={value}
                        showArrow={true}
                        style={multiSelectStyles}
                    />
                )}
            />
            <Button type="submit">Отправить</Button>
          </form>
        </div>
      </div>
  );
};

export default EditDisciplinesModal;
